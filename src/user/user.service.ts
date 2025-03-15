import { Injectable, ConflictException, UnauthorizedException, NotFoundException, BadRequestException } from '@nestjs/common';
import { type CreateUserDTO } from './dto/create-user.dto';
import { PrismaService } from 'src/prisma/prisma.service';
import { compare, hash } from 'bcrypt';
import { ConfigService } from '@nestjs/config';
import { UtilsService } from 'src/utils/utils.service';
import { RefreshUserTokenDTO } from './dto/refresh-token.dto';
import { LoginUserDTO } from './dto/login-user.dto';
import { jwtDecrypt } from 'jose';

@Injectable()
export class UserService {
    constructor(
        private readonly prismaService: PrismaService,
        private readonly configService: ConfigService,
        private readonly utilsService: UtilsService,
    ) 
    {}

    async createUser(createUserBody: CreateUserDTO) {
        const { email, password, username} = createUserBody;

        const userAlreadyExists = await this.prismaService.user.findFirst({
            where: {email},
            select: {id: true}
        });

        if (userAlreadyExists) {
            throw new ConflictException('Email já em uso')
        }

        const hashedPass = await hash(password, 10);

        const user = await this.prismaService.user.create({
            data: {
                email,
                name: username,
                password: hashedPass
            },
            select: {
                created_at: true,
                email: true,
                id: true,
                name: true,
                updated_at: true
            }
        });

        return user;
    }

    async authUser(authHeader: string) {
        return await this.utilsService.validateToken(authHeader);
    }

    async refreshToken(authHeader: string, refreshUserTokenBody: RefreshUserTokenDTO) {
        await this.utilsService.validateToken(authHeader);
        return await this.utilsService.createJWT(refreshUserTokenBody);
    }

    async loginUser(loginUserBody: LoginUserDTO) {
        const { email, password } = loginUserBody;
        
        const user = await this.prismaService.user.findFirst({
            where: {email},
            select: {
                email: true,
                id: true,
                password: true
            }
        })

        if (!user) {
            throw new NotFoundException('Usuário não encontrado');
        }

        const isPassEqual = await compare(password, user.password);
        
        if (!isPassEqual) {
            throw new BadRequestException('Email ou senha invalidos');
        }

        const jwt = await this.utilsService.createJWT({email, id: user.id})

        return jwt;
    }

    async logoutUser(authHeader: string) {
        const [, token] = authHeader.split(' ');

        if (!token) {
            throw new UnauthorizedException('Token inválido');
        }

        const secret = this.configService.get<string>('JWT_SECRET');

        const secretKey = new TextEncoder().encode(secret);

        const  { payload } = await jwtDecrypt(token, secretKey);

        const expTime = payload.exp as number;
        const currentTime = Math.floor(Date.now() / 1000);

        return;
    }
}
