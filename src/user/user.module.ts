import { Module } from "@nestjs/common";
import { UserController } from "./user.controller";
import { PrismaModule } from "src/prisma/prisma.module";
import { UserService } from "./user.service";
import { UtilsModule } from "src/utils/utils.module";
import { AppRedisModule } from "../redis/redis.module";

@Module({
    controllers: [UserController],
    imports: [PrismaModule, UtilsModule, AppRedisModule],
    providers: [UserService]
    
})
export class UserModule {}
