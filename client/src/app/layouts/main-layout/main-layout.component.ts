import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MainLayoutSidebarComponent } from './main-layout-sidebar/main-layout-sidebar.component';

@Component({
  selector: 'app-main-layout',
  imports: [RouterOutlet, MainLayoutSidebarComponent],
  templateUrl: './main-layout.component.html',
  styleUrl: './main-layout.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MainLayoutComponent {}
