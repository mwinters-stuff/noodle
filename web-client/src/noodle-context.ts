// eslint-disable-next-line import/no-extraneous-dependencies
import { createContext } from '@lit-labs/context';
import { Router } from '@vaadin/router';

import {
  NoodleApiApi,
  NoodleAuthApi,
  UserSession,
  Application,
  Tab,
  UsersApplicationItem,
} from './api/index.js';

export const userSessionContext = createContext<UserSession>('userSession');
export const noodleApiContext = createContext<NoodleApiApi>('noodleApi');
export const authApiContext = createContext<NoodleAuthApi>('authApi');

export interface DataCache {
  _applications: Application[];
  _tabs: Tab[];
  router?: Router;
  // _userApplications: UsersApplicationItem[]

  Applications(): Application[];

  // UserApplications(): UsersApplicationItem[];

  SetUserApplications(uai: UsersApplicationItem[]): void;

  Tabs(): Tab[];

  SetTabs(value: Tab[]): void;

  GetApplication(id: number): Application;
  GetTabIndex(tabId: number): number;
}

export const dataCacheContext = createContext<DataCache>('dataCache');
