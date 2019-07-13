import { Component, OnInit } from '@angular/core';
import { NotesService, Note } from '../notes.service';

@Component({
  selector: 'app-notes',
  templateUrl: './notes.component.html',
  styleUrls: ['./notes.component.css']
})
export class NotesComponent implements OnInit {

  public allNotes : Note[];
  public noteMessage : string; 
  public noteTitle : string;
  public noteRanking: string = "0";

  constructor(private noteService: NotesService) { }

  ngOnInit() {
      this.getAll();
  }

  getAll() {
    this.noteService.getNotesList().subscribe((data: Note[]) => { 
      this.allNotes = data.sort((x1, x2) => {
        if(x1.ranking > x2.ranking){
          return -1;
        }else if(x1.ranking < x2.ranking){
          return 1;
        }

        return 0
      });
    });
  }

  addNote(){
    var newNote : Note = {
        id : '', 
        title : this.noteTitle,
        message : this.noteMessage,
        ranking : Number(this.noteRanking)
    }

     
    this.noteService.addNote(newNote).subscribe(() => {
      this.getAll();
      this.noteTitle = '';
      this.noteRanking = '0';
      this.noteMessage = '';
    });

  }


  deleteTodo(todo: Note) {
    this.noteService.deleteTodo(todo).subscribe(() => {
      this.getAll();
    })
  }


}
