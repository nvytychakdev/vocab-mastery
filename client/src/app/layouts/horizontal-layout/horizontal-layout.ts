import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { HorizontalLayoutContent } from './horizontal-layout-content/horizontal-layout-content';
import { HorizontalLayoutHeader } from './horizontal-layout-header/horizontal-layout-header';

@Component({
  selector: 'app-horizontal-layout',
  imports: [HorizontalLayoutHeader, HorizontalLayoutContent, RouterOutlet],
  templateUrl: './horizontal-layout.html',
  styleUrl: './horizontal-layout.css',
})
export class HorizontalLayout {}
