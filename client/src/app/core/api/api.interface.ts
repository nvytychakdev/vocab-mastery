export type ApiUrlParams = Record<string, string>;

export type ApiListQuerySortDir = 'asc' | 'desc';
export type ApiQueryParams = Partial<{ include: string | string[] }>;
export type ApiListQueryParams = Partial<
  {
    sortBy: string;
    dir: ApiListQuerySortDir;
    limit: number;
    offset: number;
  } & ApiQueryParams
>;

export type ApiOptions<T = ApiUrlParams, Q = ApiQueryParams> = {
  params?: T;
  query?: Q;
};

export type ResponseCreate = {
  id: string;
};

export type ResponseList<T extends object> = {
  items: T[];
  total: number;
  offset: number;
  limit: number;
};

export type ResponseDelete = {
  ok: boolean;
};
