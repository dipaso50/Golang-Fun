import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class NotesService {

  constructor(private httpClient: HttpClient) {}

  getNotesList() {
    return this.httpClient.get(environment.gateway + '/getAll');
  }

  addNote(todo: Note) {
    return this.httpClient.post(environment.gateway + '/add', todo);
  }

  getById(noteId: string) {
    return this.httpClient.get(environment.gateway + '/getById/' + noteId);
  }

  deleteTodo(todo: Note) {
    return this.httpClient.delete(environment.gateway + '/delete/' + todo.id);
  }
}


export class Note {
  id: string;
  title:string
  message: string;
  ranking: number;
}
