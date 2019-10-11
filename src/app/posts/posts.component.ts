import { Component, OnInit } from '@angular/core';
import { Post } from '../Post';
import { Observable } from 'rxjs';
import { PostsService } from '../posts.service';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.less']
})
export class PostsComponent implements OnInit {
  
  posts$: Observable<Post[]>;
  loading$: Observable<boolean>;

  constructor(private postService: PostsService) { 
    this.posts$ = postService.entities$;
    this.loading$ = postService.loading$;
  }

  ngOnInit() {
  }

  getAll(){
    this.postService.getAll();
  }

  create(newPost: Post){
    let crap = new Post();
    crap.content = "contetnt";
    crap.title = 'title';
    crap.id = '1';
    this.postService.add(crap);
  }

  read(){
    this.postService.getByKey('1');
  }

  update(updatedPost: Post){
    let crap = new Post();
    crap.content = "updated";
    crap.title = 'title';
    crap.id = '1';
    this.postService.update(crap);
  }

  delete(post: Post){
    let crap = new Post();
    crap.content = "contetnt";
    crap.title = 'title';
    crap.id = '1';
    this.postService.delete(crap);
  }
}
