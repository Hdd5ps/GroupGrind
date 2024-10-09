import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { GoogleApiService } from '../services/google-api.service';

@Component({
  selector: 'app-schedule',
  templateUrl: './schedule.component.html',
  styleUrls: ['./schedule.component.css']
})
export class ScheduleComponent implements OnInit {
  studySessions: any[] = [];

  constructor(private http: HttpClient, private googleApi: GoogleApiService) {}

  ngOnInit(): void {
    this.getStudySessions();
  }

  // Fetch study sessions from the server
  getStudySessions() {
    this.http.get('/api/study-sessions').subscribe(data => {
      this.studySessions = data;
    });
  }

  // Add a new study session using Google Calendar API
  addNewSession() {
    const newSession = {
      summary: 'New Study Session',
      description: 'Description of the study session',
      start: {
        dateTime: '2024-10-10T10:00:00-04:00',
        timeZone: 'America/New_York'
      },
      end: {
        dateTime: '2024-10-10T12:00:00-04:00',
        timeZone: 'America/New_York'
      }
    };

    this.googleApi.createEvent(newSession).subscribe(response => {
      console.log('Event created:', response);
      this.getStudySessions(); // Refresh the list of study sessions
    });
  }
}
