import { HttpContext, HttpContextToken } from '@angular/common/http';

export const IS_AUTHORIZED_REQUEST = new HttpContextToken<boolean>(() => false);
export const IsAuthorizedContext = new HttpContext().set(IS_AUTHORIZED_REQUEST, true);
