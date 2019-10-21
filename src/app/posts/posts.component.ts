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
  postToEdit: Post;
  newPost: Post = new Post();

  constructor(private postService: PostsService) { 
    this.posts$ = postService.entities$;
    this.loading$ = postService.loading$;
  }

  ngOnInit() {
    this.getAll();
  }
  

  getAll(){
    this.postService.getAll();
  }

  create(newPost: Post){
    this.postService.add(newPost);
  }

  read(){
    this.postService.getByKey('1');
  }

  update(updatedPost: Post){
    this.postService.update(updatedPost);
  }

  delete(post: Post){
    this.postService.delete(post);
  }

  editPost(post: Post){
    this.postToEdit = post;
  }

  saveEdits(){
    this.update(this.postToEdit);
    this.postToEdit = undefined;
  }

  createNewPost(){
    this.create(this.newPost);
    this.newPost = new Post();
    this.getAll();
  }
}
