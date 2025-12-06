import { Component, inject } from '@angular/core';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { lucideMonitor, lucideMoon, lucideSun } from '@ng-icons/lucide';
import { Button } from '@vm/ui';
import { ThemeLoader } from '../../../../styles/theme-loader';

@Component({
  selector: 'app-horizontal-layout-header-theme-selector',
  imports: [Button, NgIcon],
  providers: [
    provideIcons({
      lucideMoon,
      lucideSun,
      lucideMonitor,
    }),
  ],
  templateUrl: './horizontal-layout-header-theme-selector.html',
  styleUrl: './horizontal-layout-header-theme-selector.css',
})
export class HorizontalLayoutHeaderThemeSelector {
  readonly themeLoader = inject(ThemeLoader);
}
