import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-study-group',
  templateUrl: './study-group.component.html',
  styleUrls: ['./study-group.component.css']
})
export class StudyGroupComponent implements OnInit {
  studyGroups: any[] = [];

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.getStudyGroups();
  }

  getStudyGroups() {
    this.http.get('/api/study-groups').subscribe(data => {
      this.studyGroups = data;
    });
  }
}
