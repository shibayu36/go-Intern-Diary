/* tslint:disable */
/* eslint-disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: GetMyDiaries
// ====================================================

export interface GetMyDiaries_visitor_diaries {
  readonly __typename: "Diary";
  readonly id: string;
  readonly name: string;
}

export interface GetMyDiaries_visitor {
  readonly __typename: "User";
  readonly id: string;
  readonly name: string;
  readonly diaries: ReadonlyArray<GetMyDiaries_visitor_diaries>;
}

export interface GetMyDiaries {
  readonly visitor: GetMyDiaries_visitor;
}
