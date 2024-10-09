package main

// StudyGroup represents a study group
type StudyGroup struct {
  ID          int    `json:"id"`
  Name        string `json:"name"`
  Description string `json:"description"`
}

// Member represents a member of a study group
type Member struct {
  ID       int    `json:"id"`
  Name     string `json:"name"`
  Email    string `json:"email"`
  GroupID  int    `json:"groupId"`
}
