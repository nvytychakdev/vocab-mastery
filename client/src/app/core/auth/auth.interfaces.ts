import { HttpContextToken } from '@angular/common/http';

export const IS_AUTHORIZED_REQUEST = new HttpContextToken<boolean>(() => false);
