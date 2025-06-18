import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { lucideChevronsUpDown, lucideGamepad, lucideHouse, lucideListPlus, lucidePanelLeft } from '@ng-icons/lucide';
import { MainLayoutSidebarProfileComponent } from './main-layout-sidebar-profile/main-layout-sidebar-profile.component';

@Component({
  selector: 'app-main-layout-sidebar',
  imports: [NgIcon, RouterLink, RouterLinkActive, MainLayoutSidebarProfileComponent],
  providers: [provideIcons({ lucideHouse, lucidePanelLeft, lucideGamepad, lucideListPlus, lucideChevronsUpDown })],
  templateUrl: './main-layout-sidebar.component.html',
  styleUrl: './main-layout-sidebar.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MainLayoutSidebarComponent {
  readonly navOptions = [
    { title: 'Home', link: '/main/home', icon: 'lucideHouse' },
    { title: 'My Words', link: '/main/my-words', icon: 'lucideListPlus' },
    { title: 'Play', link: '/main/play', icon: 'lucideGamepad' },
  ] as const;
}
