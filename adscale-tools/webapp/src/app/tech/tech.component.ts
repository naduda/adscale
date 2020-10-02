import { Component, OnInit } from '@angular/core';
import { Technology } from './tech.model';
import { TechService } from './tech.service';

@Component({
  selector: 'adscale-tech',
  templateUrl: './tech.component.html'
})
export class TechComponent implements OnInit {

  technologies: Technology[] = [];

  constructor(private readonly techService: TechService) { }

  ngOnInit() {
    this.techService.getTechnologies().subscribe(value => {
      this.technologies = value;
    });
  }
}
