import { Component } from '@angular/core';

@Component({
  selector: 'app-horizontal-layout-content',
  imports: [],
  host: {
    class: 'w-full max-w-(--width-layout) px-4',
  },
  templateUrl: './horizontal-layout-content.html',
  styleUrl: './horizontal-layout-content.css',
})
export class HorizontalLayoutContent {}
