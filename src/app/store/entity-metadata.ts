import { EntityMetadataMap } from 'ngrx-data';

const entityMetadata: EntityMetadataMap = {
  Post: {}
};


const pluralNames = { Post: 'Post' };

export const entityConfig = {
  entityMetadata,
  pluralNames
};