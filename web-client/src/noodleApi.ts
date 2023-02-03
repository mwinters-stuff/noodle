/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface Error {
  /** @format int64 */
  code?: number;
  message?: string;
  fields?: string;
}

export interface User {
  /** @format int64 */
  Id?: number;
  Username?: string;
  DN?: string;
  DisplayName?: string;
  GivenName?: string;
  Surname?: string;
  UidNumber?: number;
}

export interface Group {
  /** @format int64 */
  Id?: number;
  DN?: string;
  Name?: string;
}

export type Principal = string;

export interface UserGroup {
  /** @format int64 */
  Id?: number;
  /** @format int64 */
  GroupId?: number;
  GroupDN?: string;
  GroupName?: string;
  /** @format int64 */
  UserId?: number;
  UserDN?: string;
  UserName?: string;
}

export interface UserApplications {
  /** @format int64 */
  Id?: number;
  /** @format int64 */
  ApplicationId?: number;
  /** @format int64 */
  UserId?: number;
  Application?: Application;
}

export interface Tab {
  /** @format int64 */
  Id?: number;
  Label?: string;
  DisplayOrder?: number;
}

export interface GroupApplications {
  /** @format int64 */
  Id?: number;
  /** @format int64 */
  ApplicationId?: number;
  /** @format int64 */
  GroupId?: number;
  Application?: Application;
}

export interface Application {
  /** @format int64 */
  Id?: number;
  TemplateAppid?: string;
  Name?: string;
  Website?: string;
  License?: string;
  Description?: string;
  Enhanced?: boolean;
  TileBackground?: string;
  Icon?: string;
}

export interface ApplicationTab {
  /** @format int64 */
  Id?: number;
  /** @format int64 */
  ApplicationId?: number;
  /** @format int64 */
  TabId?: number;
  DisplayOrder?: number;
  Application?: Application;
}

export interface ApplicationTemplate {
  Appid?: string;
  Name?: string;
  Website?: string;
  License?: string;
  Description?: string;
  Enhanced?: boolean;
  tile_background?: string;
  Icon?: string;
  SHA?: string;
}

export interface AppList {
  AppCount?: number;
  Apps?: ApplicationTemplate[];
}

export interface UserSession {
  /** @format int64 */
  Id?: number;
  /** @format int64 */
  UserId?: number;
  Token?: string;
  /** @format date-time */
  Issued?: string;
  /** @format date-time */
  Expires?: string;
}

export interface UserLogin {
  username?: string;
  /** @format password */
  password?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, 'body' | 'bodyUsed'>;

export interface FullRequestParams extends Omit<RequestInit, 'body'> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<
  FullRequestParams,
  'body' | 'method' | 'query' | 'path'
>;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, 'baseUrl' | 'cancelToken' | 'signal'>;
  securityWorker?: (
    securityData: SecurityDataType | null
  ) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown>
  extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = 'application/json',
  FormData = 'multipart/form-data',
  UrlEncoded = 'application/x-www-form-urlencoded',
  Text = 'text/plain',
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = 'http://localhost:8081/api';
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>['securityWorker'];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) =>
    fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: 'same-origin',
    headers: {},
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(
      typeof value === 'number' ? value : `${value}`
    )}`;
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join('&');
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter(
      key => 'undefined' !== typeof query[key]
    );
    return keys
      .map(key =>
        Array.isArray(query[key])
          ? this.addArrayQueryParam(query, key)
          : this.addQueryParam(query, key)
      )
      .join('&');
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : '';
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === 'object' || typeof input === 'string')
        ? JSON.stringify(input)
        : input,
    [ContentType.Text]: (input: any) =>
      input !== null && typeof input !== 'string'
        ? JSON.stringify(input)
        : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === 'object' && property !== null
            ? JSON.stringify(property)
            : `${property}`
        );
        return formData;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  protected mergeRequestParams(
    params1: RequestParams,
    params2?: RequestParams
  ): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected createAbortSignal = (
    cancelToken: CancelToken
  ): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === 'boolean' ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(
      `${baseUrl || this.baseUrl || ''}${path}${
        queryString ? `?${queryString}` : ''
      }`,
      {
        ...requestParams,
        headers: {
          ...(requestParams.headers || {}),
          ...(type && type !== ContentType.FormData
            ? { 'Content-Type': type }
            : {}),
        },
        signal: cancelToken
          ? this.createAbortSignal(cancelToken)
          : requestParams.signal,
        body:
          typeof body === 'undefined' || body === null
            ? null
            : payloadFormatter(body),
      }
    ).then(async response => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then(data => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch(e => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title Noodle
 * @version 2.0
 * @license Apache License (https://gitea.winters.org.nz/mathew/noodle/src/branch/master/LICENSE)
 * @baseUrl http://localhost:8081/api
 * @contact Source Code (https://gitea.winters.org.nz/mathew/noodle)
 *
 * Noodle
 */
export class Api<
  SecurityDataType extends unknown
> extends HttpClient<SecurityDataType> {
  healthz = {
    /**
     * @description used by Kubernetes liveness probe
     *
     * @tags Kubernetes
     * @name HealthzList
     * @summary Liveness check
     * @request GET:/healthz
     */
    healthzList: (params: RequestParams = {}) =>
      this.request<object, any>({
        path: `/healthz`,
        method: 'GET',
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),
  };
  readyz = {
    /**
     * @description used by Kubernetes readiness probe
     *
     * @tags Kubernetes
     * @name ReadyzList
     * @summary Readiness check
     * @request GET:/readyz
     */
    readyzList: (params: RequestParams = {}) =>
      this.request<object, any>({
        path: `/readyz`,
        method: 'GET',
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),
  };
  auth = {
    /**
     * @description Authenticates a User
     *
     * @tags noodle-auth
     * @name AuthenticateCreate
     * @request POST:/auth/authenticate
     */
    authenticateCreate: (login: UserLogin, params: RequestParams = {}) =>
      this.request<UserSession, Error>({
        path: `/auth/authenticate`,
        method: 'POST',
        body: login,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Log out a user
     *
     * @tags noodle-auth
     * @name LogoutList
     * @request GET:/auth/logout
     * @secure
     */
    logoutList: (params: RequestParams = {}) =>
      this.request<void, Error>({
        path: `/auth/logout`,
        method: 'GET',
        secure: true,
        ...params,
      }),
  };
  noodle = {
    /**
     * @description Gets the list of users or a single user
     *
     * @tags noodle-api
     * @name UsersList
     * @request GET:/noodle/users
     * @secure
     */
    usersList: (
      query?: {
        userid?: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<User[], Error>({
        path: `/noodle/users`,
        method: 'GET',
        query: query,
        secure: true,
        format: 'json',
        ...params,
      }),

    /**
     * @description Gets the list of groups
     *
     * @tags noodle-api
     * @name GroupsList
     * @request GET:/noodle/groups
     * @secure
     */
    groupsList: (
      query?: {
        groupid?: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<Group[], Error>({
        path: `/noodle/groups`,
        method: 'GET',
        query: query,
        secure: true,
        format: 'json',
        ...params,
      }),

    /**
     * @description Gets the list of Groups for a user or users for a group
     *
     * @tags noodle-api
     * @name UserGroupsList
     * @request GET:/noodle/user-groups
     * @secure
     */
    userGroupsList: (
      query?: {
        userid?: number;
        groupid?: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<UserGroup[], Error>({
        path: `/noodle/user-groups`,
        method: 'GET',
        query: query,
        secure: true,
        format: 'json',
        ...params,
      }),

    /**
     * @description Loads Users and Groups to Database
     *
     * @tags noodle-api
     * @name LdapReloadList
     * @request GET:/noodle/ldap/reload
     * @secure
     */
    ldapReloadList: (params: RequestParams = {}) =>
      this.request<void, Error>({
        path: `/noodle/ldap/reload`,
        method: 'GET',
        secure: true,
        ...params,
      }),

    /**
     * @description Loads Hiemdall App Templates to Database
     *
     * @tags noodle-api
     * @name HeimdallReloadList
     * @request GET:/noodle/heimdall/reload
     * @secure
     */
    heimdallReloadList: (params: RequestParams = {}) =>
      this.request<void, Error>({
        path: `/noodle/heimdall/reload`,
        method: 'GET',
        secure: true,
        ...params,
      }),

    /**
     * @description Gets the list of tabs
     *
     * @tags noodle-api
     * @name TabsList
     * @request GET:/noodle/tabs
     * @secure
     */
    tabsList: (params: RequestParams = {}) =>
      this.request<Tab[], Error>({
        path: `/noodle/tabs`,
        method: 'GET',
        secure: true,
        format: 'json',
        ...params,
      }),

    /**
     * @description Adds a new tab
     *
     * @tags noodle-api
     * @name TabsCreate
     * @request POST:/noodle/tabs
     * @secure
     */
    tabsCreate: (
      query: {
        action: 'insert' | 'update';
      },
      tab: Tab,
      params: RequestParams = {}
    ) =>
      this.request<Tab, Error>({
        path: `/noodle/tabs`,
        method: 'POST',
        query: query,
        body: tab,
        secure: true,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Deletes the tab
     *
     * @tags noodle-api
     * @name TabsDelete
     * @request DELETE:/noodle/tabs
     * @secure
     */
    tabsDelete: (
      query: {
        tabid: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<void, Error>({
        path: `/noodle/tabs`,
        method: 'DELETE',
        query: query,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Gets the list of applications under the tab
     *
     * @tags noodle-api
     * @name ApplicationTabsList
     * @request GET:/noodle/application-tabs
     * @secure
     */
    applicationTabsList: (
      query: {
        tab_id: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<ApplicationTab[], Error>({
        path: `/noodle/application-tabs`,
        method: 'GET',
        query: query,
        secure: true,
        format: 'json',
        ...params,
      }),

    /**
     * @description Adds a new application in a  tab
     *
     * @tags noodle-api
     * @name ApplicationTabsCreate
     * @request POST:/noodle/application-tabs
     * @secure
     */
    applicationTabsCreate: (
      query: {
        action: 'insert' | 'update';
      },
      application_tab: ApplicationTab,
      params: RequestParams = {}
    ) =>
      this.request<ApplicationTab, Error>({
        path: `/noodle/application-tabs`,
        method: 'POST',
        query: query,
        body: application_tab,
        secure: true,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Deletes the application_tab
     *
     * @tags noodle-api
     * @name ApplicationTabsDelete
     * @request DELETE:/noodle/application-tabs
     * @secure
     */
    applicationTabsDelete: (
      query: {
        application_tab_id: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<void, Error>({
        path: `/noodle/application-tabs`,
        method: 'DELETE',
        query: query,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Gets the list of user applications
     *
     * @tags noodle-api
     * @name UserApplicationsList
     * @request GET:/noodle/user-applications
     * @secure
     */
    userApplicationsList: (
      query: {
        user_id: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<UserApplications[], Error>({
        path: `/noodle/user-applications`,
        method: 'GET',
        query: query,
        secure: true,
        format: 'json',
        ...params,
      }),

    /**
     * @description Adds a new user application
     *
     * @tags noodle-api
     * @name UserApplicationsCreate
     * @request POST:/noodle/user-applications
     * @secure
     */
    userApplicationsCreate: (
      user_application: UserApplications,
      params: RequestParams = {}
    ) =>
      this.request<UserApplications, Error>({
        path: `/noodle/user-applications`,
        method: 'POST',
        body: user_application,
        secure: true,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Deletes the user application
     *
     * @tags noodle-api
     * @name UserApplicationsDelete
     * @request DELETE:/noodle/user-applications
     * @secure
     */
    userApplicationsDelete: (
      query: {
        user_application_id: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<void, Error>({
        path: `/noodle/user-applications`,
        method: 'DELETE',
        query: query,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Gets the list of group applications
     *
     * @tags noodle-api
     * @name GroupApplicationsList
     * @request GET:/noodle/group-applications
     * @secure
     */
    groupApplicationsList: (
      query: {
        group_id: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<GroupApplications[], Error>({
        path: `/noodle/group-applications`,
        method: 'GET',
        query: query,
        secure: true,
        format: 'json',
        ...params,
      }),

    /**
     * @description Adds a new group application
     *
     * @tags noodle-api
     * @name GroupApplicationsCreate
     * @request POST:/noodle/group-applications
     * @secure
     */
    groupApplicationsCreate: (
      group_application: GroupApplications,
      params: RequestParams = {}
    ) =>
      this.request<GroupApplications, Error>({
        path: `/noodle/group-applications`,
        method: 'POST',
        body: group_application,
        secure: true,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Deletes the group application
     *
     * @tags noodle-api
     * @name GroupApplicationsDelete
     * @request DELETE:/noodle/group-applications
     * @secure
     */
    groupApplicationsDelete: (
      query: {
        group_application_id: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<void, Error>({
        path: `/noodle/group-applications`,
        method: 'DELETE',
        query: query,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Gets the list of application templates
     *
     * @tags noodle-api
     * @name AppTemplatesList
     * @request GET:/noodle/app-templates
     * @secure
     */
    appTemplatesList: (
      query: {
        search: string;
      },
      params: RequestParams = {}
    ) =>
      this.request<ApplicationTemplate[], Error>({
        path: `/noodle/app-templates`,
        method: 'GET',
        query: query,
        secure: true,
        format: 'json',
        ...params,
      }),

    /**
     * @description Gets application by id or template_id
     *
     * @tags noodle-api
     * @name ApplicationsList
     * @request GET:/noodle/applications
     * @secure
     */
    applicationsList: (
      query?: {
        application_id?: number;
        application_template?: string;
      },
      params: RequestParams = {}
    ) =>
      this.request<Application[], Error>({
        path: `/noodle/applications`,
        method: 'GET',
        query: query,
        secure: true,
        format: 'json',
        ...params,
      }),

    /**
     * @description Adds a new application
     *
     * @tags noodle-api
     * @name ApplicationsCreate
     * @request POST:/noodle/applications
     * @secure
     */
    applicationsCreate: (
      application: Application,
      params: RequestParams = {}
    ) =>
      this.request<Application, Error>({
        path: `/noodle/applications`,
        method: 'POST',
        body: application,
        secure: true,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * @description Deletes the application
     *
     * @tags noodle-api
     * @name ApplicationsDelete
     * @request DELETE:/noodle/applications
     * @secure
     */
    applicationsDelete: (
      query: {
        application_id: number;
      },
      params: RequestParams = {}
    ) =>
      this.request<void, Error>({
        path: `/noodle/applications`,
        method: 'DELETE',
        query: query,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),
  };
}
