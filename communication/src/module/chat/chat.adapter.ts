import { IoAdapter } from '@nestjs/platform-socket.io';
import Redis from 'ioredis';
import { ServerOptions } from 'socket.io';
import { createAdapter } from '@socket.io/redis-adapter';
import { NestFastifyApplication } from '@nestjs/platform-fastify';

export class ChatAdapter extends IoAdapter {
  private adapterConstructor: ReturnType<typeof createAdapter>;

  constructor(app: NestFastifyApplication, private readonly redis: Redis) {
    super(app);
  }

  async connectToRedis(): Promise<void> {
    this.adapterConstructor = createAdapter(this.redis, this.redis);
  }

  createIOServer(port: number, options?: ServerOptions) {
    const server = super.createIOServer(port, options);
    server.adapter(this.adapterConstructor);
    return server;
  }
}
