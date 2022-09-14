import { Catch, ArgumentsHost, BadRequestException } from '@nestjs/common';
import { BaseWsExceptionFilter } from '@nestjs/websockets';
import { Socket } from 'socket.io';
import { EmitEvent } from 'src/enum/event.enum';
import { EmitEvents } from 'src/type/event.type';

@Catch()
export class AllExceptionsFilter extends BaseWsExceptionFilter {
  catch(exception: Error, host: ArgumentsHost) {
    const client = host.switchToWs().getClient() as Socket<null, EmitEvents>;
    let errorResponse: string | string[];

    if (exception instanceof BadRequestException) {
      const res = exception.getResponse();
      errorResponse =
        typeof res === 'string' ? res : (res as any)?.message?.[0];
    } else {
      errorResponse = exception.message;
    }

    client.emit(EmitEvent.Error, {
      event: client.eventName,
      message: errorResponse,
    });
  }
}
