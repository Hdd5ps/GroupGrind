import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit {
  user: any = {};
  userStudyGroups: any[] = [];

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.getUserInfo();
    this.getUserStudyGroups();
  }

  // Fetch user information from the server
  getUserInfo() {
    this.http.get('/api/user').subscribe(data => {
      this.user = data;
    });
  }

  // Fetch user's study groups from the server
  getUserStudyGroups() {
    this.http.get('/api/user/study-groups').subscribe(data => {
      this.userStudyGroups = data;
    });
  }

  // Update user details
  updateUser() {
    this.http.put('/api/user', this.user).subscribe(response => {
      console.log('User updated successfully');
    });
  }
}
