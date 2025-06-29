export type Entity<T extends Record<string, unknown>> = {
  id: string;
  craetedAt: string;
} & T;

export type StripEntity<T> = T extends Entity<infer U> ? U : never;
