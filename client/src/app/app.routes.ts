import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, Router, Routes } from '@angular/router';
import { catchError, map, of } from 'rxjs';
import { authProfileResolve } from './features/auth/auth-profile.resolver';
import { AuthService } from './features/auth/auth.service';
import { authRedirectGuards } from './features/auth/guards/auth.guard';
import { HorizontalLayout } from './layouts/horizontal-layout/horizontal-layout';
import { ConfirmEmailComponent } from './pages/auth/confirm-email/confirm-email.component';
import { SignInComponent } from './pages/auth/sign-in/sign-in.component';
import { SignUpComponent } from './pages/auth/sign-up/sign-up.component';
import { UsedEmailComponent } from './pages/auth/used-email/used-email.component';
import { HomeComponent } from './pages/home/home.component';
import { MyWords } from './pages/my-words/my-words';
import { PlayComponent } from './pages/play/play.component';

const { redirectIfAuthenticated, redirectIfUnauthenticated } = authRedirectGuards({
  redirectAuth: '/main',
  redirectUnauth: '/auth',
});

const confirmEmailRedirect = (route: ActivatedRouteSnapshot) => {
  const router = inject(Router);
  const auth = inject(AuthService);

  const token = route.queryParamMap.get('token');
  if (token) {
    return auth.confirmEmail(token).pipe(
      map(() => router.createUrlTree(['/main'])),
      catchError(() => {
        return of(router.createUrlTree(['/auth/used-email']));
      })
    );
  }

  return true;
};

export const routes: Routes = [
  {
    path: 'auth',
    canActivate: [redirectIfAuthenticated],
    children: [
      {
        path: 'sign-in',
        component: SignInComponent,
      },
      {
        path: 'sign-up',
        component: SignUpComponent,
      },
      {
        path: 'confirm-email',
        canActivate: [confirmEmailRedirect],
        component: ConfirmEmailComponent,
      },
      {
        path: 'used-email',
        component: UsedEmailComponent,
      },
      {
        path: '**',
        redirectTo: 'sign-in',
      },
    ],
  },
  {
    path: 'main',
    component: HorizontalLayout,
    canActivate: [redirectIfUnauthenticated],
    resolve: [authProfileResolve],
    children: [
      {
        path: 'home',
        component: HomeComponent,
      },
      {
        path: 'my-words',
        children: [
          {
            path: '',
            component: MyWords,
          },
        ],
      },
      {
        path: 'play',
        component: PlayComponent,
      },
      {
        path: '**',
        redirectTo: 'home',
      },
    ],
  },
  {
    path: '**',
    redirectTo: 'auth',
  },
];
