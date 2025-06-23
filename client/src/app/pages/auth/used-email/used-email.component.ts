import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { Button } from '@vm/ui';
import { FullscreenLayoutComponent } from '../../../layouts/fullscreen-layout/fullscreen-layout.component';

@Component({
  selector: 'app-used-email',
  imports: [FullscreenLayoutComponent, RouterLink, Button],
  templateUrl: './used-email.component.html',
  styleUrl: './used-email.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UsedEmailComponent {}
