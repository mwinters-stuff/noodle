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
 * @interface ApplicationTemplate
 */
export interface ApplicationTemplate {
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  appid?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  name?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  website?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  license?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  description?: string;
  /**
   *
   * @type {boolean}
   * @memberof ApplicationTemplate
   */
  enhanced?: boolean;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  tileBackground?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  icon?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  sHA?: string;
}

/**
 * Check if a given object implements the ApplicationTemplate interface.
 */
export function instanceOfApplicationTemplate(value: object): boolean {
  let isInstance = true;

  return isInstance;
}

export function ApplicationTemplateFromJSON(json: any): ApplicationTemplate {
  return ApplicationTemplateFromJSONTyped(json, false);
}

export function ApplicationTemplateFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean
): ApplicationTemplate {
  if (json === undefined || json === null) {
    return json;
  }
  return {
    appid: !exists(json, 'Appid') ? undefined : json['Appid'],
    name: !exists(json, 'Name') ? undefined : json['Name'],
    website: !exists(json, 'Website') ? undefined : json['Website'],
    license: !exists(json, 'License') ? undefined : json['License'],
    description: !exists(json, 'Description') ? undefined : json['Description'],
    enhanced: !exists(json, 'Enhanced') ? undefined : json['Enhanced'],
    tileBackground: !exists(json, 'tile_background')
      ? undefined
      : json['tile_background'],
    icon: !exists(json, 'Icon') ? undefined : json['Icon'],
    sHA: !exists(json, 'SHA') ? undefined : json['SHA'],
  };
}

export function ApplicationTemplateToJSON(
  value?: ApplicationTemplate | null
): any {
  if (value === undefined) {
    return undefined;
  }
  if (value === null) {
    return null;
  }
  return {
    Appid: value.appid,
    Name: value.name,
    Website: value.website,
    License: value.license,
    Description: value.description,
    Enhanced: value.enhanced,
    tile_background: value.tileBackground,
    Icon: value.icon,
    SHA: value.sHA,
  };
}
