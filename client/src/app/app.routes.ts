import { Routes } from '@angular/router';
import { MainLayoutComponent } from './layouts/main-layout/main-layout.component';
import { HomeComponent } from './pages/home/home.component';
import { MyWordsComponent } from './pages/my-words/my-words.component';
import { PlayComponent } from './pages/play/play.component';

export const routes: Routes = [
  {
    path: '',
    component: MainLayoutComponent,
    children: [
      {
        path: 'home',
        component: HomeComponent,
      },
      {
        path: 'my-words',
        component: MyWordsComponent,
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
];
