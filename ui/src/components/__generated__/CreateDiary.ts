/* tslint:disable */
/* eslint-disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: CreateDiary
// ====================================================

export interface CreateDiary_createDiary {
  readonly __typename: "Diary";
  readonly id: string;
  readonly name: string;
}

export interface CreateDiary {
  readonly createDiary: CreateDiary_createDiary;
}

export interface CreateDiaryVariables {
  readonly name: string;
}
