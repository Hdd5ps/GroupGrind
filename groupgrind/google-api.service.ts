import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class GoogleApiService {
  private apiUrl = 'https://www.googleapis.com/calendar/v3/calendars/primary/events';
  private accessToken = 'YOUR_ACCESS_TOKEN'; // Replace with your actual access token

  constructor(private http: HttpClient) {}

  // Create a new event in Google Calendar
  createEvent(event: any): Observable<any> {
    const headers = new HttpHeaders({
      'Authorization': `Bearer ${this.accessToken}`,
      'Content-Type': 'application/json'
    });

    return this.http.post(this.apiUrl, event, { headers });
  }
}
