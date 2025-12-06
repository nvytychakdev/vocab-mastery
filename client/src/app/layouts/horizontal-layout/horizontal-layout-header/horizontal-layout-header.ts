import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { lucideBell } from '@ng-icons/lucide';
import { Button } from '@vm/ui';
import { HorizontalLayoutHeaderProfile } from './horizontal-layout-header-profile/horizontal-layout-header-profile';
import { HorizontalLayoutHeaderThemeSelector } from './horizontal-layout-header-theme-selector/horizontal-layout-header-theme-selector';

@Component({
  selector: 'app-horizontal-layout-header',
  imports: [
    Button,
    RouterLink,
    HorizontalLayoutHeaderProfile,
    HorizontalLayoutHeaderThemeSelector,
    RouterLinkActive,
    NgIcon,
  ],
  host: {
    class: 'flex items-center justify-between gap-2',
  },
  providers: [provideIcons({ lucideBell })],
  templateUrl: './horizontal-layout-header.html',
  styleUrl: './horizontal-layout-header.css',
})
export class HorizontalLayoutHeader {
  readonly navOptions = [
    { title: 'Home', link: '/main/home', icon: 'lucideHouse' },
    { title: 'Words', link: '/main/my-words', icon: 'lucideListPlus' },
    { title: 'Flashcards', link: '/main/play', icon: 'lucideGamepad' },
  ] as const;
}
