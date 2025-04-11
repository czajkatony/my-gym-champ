import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { CarouselModule } from 'primeng/carousel';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, RouterModule, CardModule, ButtonModule, CarouselModule],
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent {
  features = [
    {
      title: 'Digital Membership',
      description: 'Get a digital membership card for quick access to your gym.',
      icon: 'pi pi-id-card'
    },
    {
      title: 'Smart Lock Access',
      description: 'Easy access to your gym with smart lock technology.',
      icon: 'pi pi-lock'
    },
    {
      title: 'Workout Tracking',
      description: 'Track your workouts, sets, reps, and personal records.',
      icon: 'pi pi-chart-line'
    },
    {
      title: 'Community Leaderboard',
      description: 'Compete with fellow gym members and track your progress.',
      icon: 'pi pi-chart-bar'
    },
    {
      title: 'Gym Events',
      description: 'Stay updated on gym events, competitions, and promotions.',
      icon: 'pi pi-calendar'
    }
  ];
} 