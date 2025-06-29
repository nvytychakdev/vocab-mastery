import { Entity } from '@core/interfaces/dto.interface';

export type Profile = Entity<{
  name: string;
  email: string;
  pictureUrl?: string;
}>;
