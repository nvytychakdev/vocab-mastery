import { Routes } from '@angular/router';
import { authProfileResolve } from './core/auth/auth-profile.resolver';
import { authRedirectGuards } from './core/auth/guards/auth.guard';
import { MainLayoutComponent } from './layouts/main-layout/main-layout.component';
import { SignInComponent } from './pages/auth/sign-in/sign-in.component';
import { SignUpComponent } from './pages/auth/sign-up/sign-up.component';
import { HomeComponent } from './pages/home/home.component';
import { MyWordsDictionaryComponent } from './pages/my-words/my-words-dictionary/my-words-dictionary.component';
import { MyWordsComponent } from './pages/my-words/my-words.component';
import { PlayComponent } from './pages/play/play.component';

const { redirectIfAuthenticated, redirectIfUnauthenticated } = authRedirectGuards({
  redirectAuth: '/main',
  redirectUnauth: '/auth',
});

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
        path: '**',
        redirectTo: 'sign-in',
      },
    ],
  },
  {
    path: 'main',
    component: MainLayoutComponent,
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
            component: MyWordsComponent,
          },
          {
            path: ':dictionaryId',
            component: MyWordsDictionaryComponent,
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
