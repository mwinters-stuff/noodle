// eslint-disable-next-line import/no-extraneous-dependencies
import { createContext } from '@lit-labs/context';

import { NoodleApiApi, NoodleAuthApi, UserSession } from './api/index.js';

export const userSessionContext = createContext<UserSession>('userSession');
export const noodleApiContext = createContext<NoodleApiApi>('noodleApi');
export const authApiContext = createContext<NoodleAuthApi>('authApi');
