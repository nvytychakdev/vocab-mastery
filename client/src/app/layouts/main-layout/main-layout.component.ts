import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { lucideGamepad, lucideHouse, lucideListPlus } from '@ng-icons/lucide';

@Component({
  selector: 'app-main-layout',
  imports: [RouterOutlet, NgIcon, RouterLink, RouterLinkActive],
  providers: [provideIcons({ lucideHouse, lucideGamepad, lucideListPlus })],
  templateUrl: './main-layout.component.html',
  styleUrl: './main-layout.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MainLayoutComponent {
  readonly navOptions = [
    { title: 'Home', link: '/main/home', icon: 'lucideHouse' },
    { title: 'My Words', link: '/main/my-words', icon: 'lucideListPlus' },
    { title: 'Play', link: '/main/play', icon: 'lucideGamepad' },
  ] as const;
}
