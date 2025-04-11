import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MenubarModule } from 'primeng/menubar';
import { ButtonModule } from 'primeng/button';
import { MenuItem } from 'primeng/api';

@Component({
  standalone: true,
  imports: [RouterModule, MenubarModule, ButtonModule],
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {
  title = 'My Gym Champ';
  
  items: MenuItem[] = [
    {
      label: 'Home',
      icon: 'pi pi-home',
      routerLink: '/'
    },
    {
      label: 'Gyms',
      icon: 'pi pi-building',
      routerLink: '/gyms'
    },
    {
      label: 'Workouts',
      icon: 'pi pi-chart-bar',
      routerLink: '/workouts'
    },
    {
      label: 'Events',
      icon: 'pi pi-calendar',
      routerLink: '/events'
    }
  ];
}
