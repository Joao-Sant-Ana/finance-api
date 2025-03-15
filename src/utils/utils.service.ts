import { Injectable, InternalServerErrorException, UnauthorizedException } from "@nestjs/common";
import { ConfigService } from "@nestjs/config";
import { jwtVerify, SignJWT } from "jose";

@Injectable()
export class UtilsService {
    constructor(
        private readonly configService: ConfigService,
    ) {}
    
    async validateToken(authHeader: string) {
        const [type, token] = authHeader.split(' ');

        if (type !== 'Bearer') {
            throw new UnauthorizedException('Token deve ser um Bearer');
        }

        if (!token) {
            throw new UnauthorizedException('Token inválido');
        }

        const secret = this.configService.get<string>('JWT_SECRET');

        if (!secret) {
            throw new InternalServerErrorException('Erro interno do servidor');
        }

        const secretKey = new TextEncoder().encode(secret);

        const { payload } = await jwtVerify(token, secretKey);

        const expectedIssuer = this.configService.get<string>('JWT_ISSUER');

        if (expectedIssuer && payload.iss !== expectedIssuer) {
            throw new UnauthorizedException('Token inválido');
        }

        return;
    }

    async createJWT(data: {email: string, id: string}) {
        const secret = this.configService.get<string>('JWT_SECRET');
        const issuer = this.configService.get<string>('JWT_ISSUER');

        const secretKey = new TextEncoder().encode(secret);

        const { id, email } = data;

        const jwt = new SignJWT({id, email})
            .setProtectedHeader({alg: 'HS256'})
            .setIssuedAt()
            .setExpirationTime('7d')
            .setIssuer(issuer)
            .sign(secretKey)


        return jwt;
    }
}