import { Body, Controller, HttpCode, HttpStatus, Post, Headers, Get, Response, Put } from '@nestjs/common';
import { Response as ExpressRespnse } from 'express';
import { UserService } from './user.service';
import { CreateUserDTO } from './dto/create-user.dto';
import { RefreshUserTokenDTO } from './dto/refresh-token.dto';
import { LoginUserDTO } from './dto/login-user.dto';

@Controller('user')
export class UserController {
  constructor(private readonly userService: UserService) {}

    @Post()
    @HttpCode(HttpStatus.CREATED)
    async createUser(@Body() createUserBody: CreateUserDTO) {
        return await this.userService.createUser(createUserBody);
    }

    @Get('auth')
    @HttpCode(HttpStatus.NO_CONTENT)
    async authUser(@Headers('Authorization') authHeader: string) {
      return await this.userService.authUser(authHeader)
    }

    @Post('refresh-token')
    @HttpCode(HttpStatus.OK)
    async refreshToken(@Headers('Authorization') authHeader: string, @Body() refreshUserTokenBody: RefreshUserTokenDTO, @Response({ passthrough: true }) response: ExpressRespnse) {
      const jwt = await this.userService.refreshToken(authHeader, refreshUserTokenBody);
      response.cookie('token', jwt, {
        httpOnly: true,
        secure: false, //Use true in prod
        sameSite: 'strict',
        maxAge: 7 * 24 * 60 * 60 * 1000
      })
    
      return { success: true};
    }

    @Post('login')
    @HttpCode(HttpStatus.OK)
    async loginUser(@Body() loginUserBody: LoginUserDTO, @Response({ passthrough: true }) response: ExpressRespnse) {
      const jwt  = await this.userService.loginUser(loginUserBody);

      response.cookie('token', jwt, {
        httpOnly: true,
        secure: false, //Use true in prod
        sameSite: 'strict',
        maxAge: 7 * 24 * 60 * 60 * 1000
      })

      return { success: true}
    }

    @Post('logout')
    @HttpCode(HttpStatus.OK)
    async logoutUser(@Headers('Authorization') authHeader: string, @Response({ passthrough: true }) response: ExpressRespnse) {
      await this.userService.logoutUser(authHeader);

      response.clearCookie('token');
      return { success: true};
    }
}
