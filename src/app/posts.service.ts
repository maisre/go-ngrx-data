import { Injectable } from '@angular/core';
import {
  EntityCollectionServiceBase,
  EntityCollectionServiceElementsFactory
} from 'ngrx-data';
import { Post } from './Post';

@Injectable({
  providedIn: 'root'
})
export class PostsService extends EntityCollectionServiceBase<Post> {

  constructor(serviceElementsFactory: EntityCollectionServiceElementsFactory) {
    super('Post', serviceElementsFactory);
   }
}
