export type Entity<T extends Record<string, unknown>> = {
  id: string;
  createdAt: string;
} & T;

export type StripEntity<T> = T extends Entity<infer U> ? U : never;
