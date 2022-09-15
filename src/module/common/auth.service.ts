import { Injectable } from '@nestjs/common';
import { User } from '@prisma/client';
import { Auth } from 'firebase-admin/auth';
import { FirebaseAuth } from 'src/decorator/firebase-auth.decorator';
import { UserId } from 'src/enum/user.enum';
import { PrismaService } from './prisma.service';

@Injectable()
export class AuthService {
  @FirebaseAuth()
  private readonly auth: Auth;

  constructor(private prismaService: PrismaService) {}

  /**
   * Get uid generated by firebase authentication using
   * ID token provided by it.
   *
   * @param token
   * @returns
   */
  private async getFirebaseId(token: string) {
    let fid: string;

    try {
      const decodedToken = await this.auth.verifyIdToken(token);
      fid = decodedToken.uid;
    } catch {
      fid = '';
    }

    return fid;
  }

  /**
   * Create an empty user with only id field. This function
   * is only used when an invalid user is returned.
   *
   * @param id
   * @returns
   */
  private generateEmptyUser(id: number): User {
    return {
      id,
    } as User;
  }

  /**
   * Get a corresponding user on the entered token. Return an empty
   * user if authentication failed.
   *
   * @param token ID token provided by firebase authentication.
   * @returns
   */
  async getUser(token: string) {
    const fid = await this.getFirebaseId(token);

    if (fid === '') {
      return this.generateEmptyUser(UserId.NonExist);
    }

    const user = await this.prismaService.user.findUnique({
      where: {
        fid,
      },
    });

    return user ?? this.generateEmptyUser(UserId.Asynchronous);
  }
}