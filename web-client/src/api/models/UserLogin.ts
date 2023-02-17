/* tslint:disable */
/* eslint-disable */
/**
 * Noodle
 * Noodle
 *
 * The version of the OpenAPI document: 2.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 *
 * @export
 * @interface UserLogin
 */
export interface UserLogin {
  /**
   *
   * @type {string}
   * @memberof UserLogin
   */
  username?: string;
  /**
   *
   * @type {string}
   * @memberof UserLogin
   */
  password?: string;
}

/**
 * Check if a given object implements the UserLogin interface.
 */
export function instanceOfUserLogin(value: object): boolean {
  let isInstance = true;

  return isInstance;
}

export function UserLoginFromJSON(json: any): UserLogin {
  return UserLoginFromJSONTyped(json, false);
}

export function UserLoginFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean
): UserLogin {
  if (json === undefined || json === null) {
    return json;
  }
  return {
    username: !exists(json, 'username') ? undefined : json['username'],
    password: !exists(json, 'password') ? undefined : json['password'],
  };
}

export function UserLoginToJSON(value?: UserLogin | null): any {
  if (value === undefined) {
    return undefined;
  }
  if (value === null) {
    return null;
  }
  return {
    username: value.username,
    password: value.password,
  };
}