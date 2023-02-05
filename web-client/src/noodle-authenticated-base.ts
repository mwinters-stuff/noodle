import { LitElement } from 'lit';
// eslint-disable-next-line import/no-extraneous-dependencies
import { Router } from '@vaadin/router';

import { Api } from './noodleApi.js';

import './noodle-login.js';
import './noodle-dash.js';

export abstract class NoodleAuthenticatedBase extends LitElement {
  private static authHeader(token: string) {
    return { 'X-Token': token };
  }

  private static noodleApi = new Api({
    securityWorker: token => ({
      headers: NoodleAuthenticatedBase.authHeader(token!.toString()),
    }),
  });

  public static getAuthenticatedApi() {
    const authToken = NoodleAuthenticatedBase.getAuthToken();
    NoodleAuthenticatedBase.noodleApi.setSecurityData(authToken);
    return NoodleAuthenticatedBase.noodleApi;
  }

  private static getAuthToken() {
    return document.cookie
      .split('; ')
      .find(row => row.startsWith('noodle-auth='))
      ?.split('=')[1];
  }

  firstUpdated() {
    if (!NoodleAuthenticatedBase.checkACookieExists('noodle-auth')) {
      Router.go(`/login`);
    }
  }

  public static IsAuthenticated() {
    return NoodleAuthenticatedBase.checkACookieExists('noodle-auth');
  }

  private static checkACookieExists(cookieName: string) {
    if (
      document.cookie
        .split(';')
        .some(item => item.trim().startsWith(`${cookieName}=`))
    ) {
      return true;
    }
    return false;
  }
}
