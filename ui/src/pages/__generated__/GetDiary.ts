/* tslint:disable */
/* eslint-disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: GetDiary
// ====================================================

export interface GetDiary_diary {
  readonly __typename: "Diary";
  readonly id: string;
  readonly name: string;
}

export interface GetDiary {
  readonly diary: GetDiary_diary;
}

export interface GetDiaryVariables {
  readonly diaryId: string;
}
