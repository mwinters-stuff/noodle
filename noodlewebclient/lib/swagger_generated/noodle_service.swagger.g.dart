// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'noodle_service.swagger.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Error _$ErrorFromJson(Map<String, dynamic> json) => Error(
      code: json['code'] as num?,
      message: json['message'] as String?,
      fields: json['fields'] as String?,
    );

Map<String, dynamic> _$ErrorToJson(Error instance) => <String, dynamic>{
      'code': instance.code,
      'message': instance.message,
      'fields': instance.fields,
    };

User _$UserFromJson(Map<String, dynamic> json) => User(
      id: json['Id'] as num?,
      username: json['Username'] as String?,
      dn: json['DN'] as String?,
      displayName: json['DisplayName'] as String?,
      givenName: json['GivenName'] as String?,
      surname: json['Surname'] as String?,
      uidNumber: json['UidNumber'] as int?,
    );

Map<String, dynamic> _$UserToJson(User instance) => <String, dynamic>{
      'Id': instance.id,
      'Username': instance.username,
      'DN': instance.dn,
      'DisplayName': instance.displayName,
      'GivenName': instance.givenName,
      'Surname': instance.surname,
      'UidNumber': instance.uidNumber,
    };

Group _$GroupFromJson(Map<String, dynamic> json) => Group(
      id: json['Id'] as num?,
      dn: json['DN'] as String?,
      name: json['Name'] as String?,
    );

Map<String, dynamic> _$GroupToJson(Group instance) => <String, dynamic>{
      'Id': instance.id,
      'DN': instance.dn,
      'Name': instance.name,
    };

UserGroup _$UserGroupFromJson(Map<String, dynamic> json) => UserGroup(
      id: json['Id'] as num?,
      groupId: json['GroupId'] as num?,
      groupDN: json['GroupDN'] as String?,
      groupName: json['GroupName'] as String?,
      userId: json['UserId'] as num?,
      userDN: json['UserDN'] as String?,
      userName: json['UserName'] as String?,
    );

Map<String, dynamic> _$UserGroupToJson(UserGroup instance) => <String, dynamic>{
      'Id': instance.id,
      'GroupId': instance.groupId,
      'GroupDN': instance.groupDN,
      'GroupName': instance.groupName,
      'UserId': instance.userId,
      'UserDN': instance.userDN,
      'UserName': instance.userName,
    };

UserApplications _$UserApplicationsFromJson(Map<String, dynamic> json) =>
    UserApplications(
      id: json['Id'] as num?,
      applicationId: json['ApplicationId'] as num?,
      userId: json['UserId'] as num?,
      application: json['Application'] == null
          ? null
          : Application.fromJson(json['Application'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$UserApplicationsToJson(UserApplications instance) =>
    <String, dynamic>{
      'Id': instance.id,
      'ApplicationId': instance.applicationId,
      'UserId': instance.userId,
      'Application': instance.application?.toJson(),
    };

Tab _$TabFromJson(Map<String, dynamic> json) => Tab(
      id: json['Id'] as num?,
      label: json['Label'] as String?,
      displayOrder: json['DisplayOrder'] as int?,
    );

Map<String, dynamic> _$TabToJson(Tab instance) => <String, dynamic>{
      'Id': instance.id,
      'Label': instance.label,
      'DisplayOrder': instance.displayOrder,
    };

GroupApplications _$GroupApplicationsFromJson(Map<String, dynamic> json) =>
    GroupApplications(
      id: json['Id'] as num?,
      applicationId: json['ApplicationId'] as num?,
      groupId: json['GroupId'] as num?,
      application: json['Application'] == null
          ? null
          : Application.fromJson(json['Application'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$GroupApplicationsToJson(GroupApplications instance) =>
    <String, dynamic>{
      'Id': instance.id,
      'ApplicationId': instance.applicationId,
      'GroupId': instance.groupId,
      'Application': instance.application?.toJson(),
    };

Application _$ApplicationFromJson(Map<String, dynamic> json) => Application(
      id: json['Id'] as num?,
      templateAppid: json['TemplateAppid'] as String?,
      name: json['Name'] as String?,
      website: json['Website'] as String?,
      license: json['License'] as String?,
      description: json['Description'] as String?,
      enhanced: json['Enhanced'] as bool?,
      tileBackground: json['TileBackground'] as String?,
      icon: json['Icon'] as String?,
    );

Map<String, dynamic> _$ApplicationToJson(Application instance) =>
    <String, dynamic>{
      'Id': instance.id,
      'TemplateAppid': instance.templateAppid,
      'Name': instance.name,
      'Website': instance.website,
      'License': instance.license,
      'Description': instance.description,
      'Enhanced': instance.enhanced,
      'TileBackground': instance.tileBackground,
      'Icon': instance.icon,
    };

ApplicationTab _$ApplicationTabFromJson(Map<String, dynamic> json) =>
    ApplicationTab(
      id: json['Id'] as num?,
      applicationId: json['ApplicationId'] as num?,
      tabId: json['TabId'] as num?,
      displayOrder: json['DisplayOrder'] as int?,
      application: json['Application'] == null
          ? null
          : Application.fromJson(json['Application'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$ApplicationTabToJson(ApplicationTab instance) =>
    <String, dynamic>{
      'Id': instance.id,
      'ApplicationId': instance.applicationId,
      'TabId': instance.tabId,
      'DisplayOrder': instance.displayOrder,
      'Application': instance.application?.toJson(),
    };

ApplicationTemplate _$ApplicationTemplateFromJson(Map<String, dynamic> json) =>
    ApplicationTemplate(
      appid: json['Appid'] as String?,
      name: json['Name'] as String?,
      website: json['Website'] as String?,
      license: json['License'] as String?,
      description: json['Description'] as String?,
      enhanced: json['Enhanced'] as bool?,
      tileBackground: json['tile_background'] as String?,
      icon: json['Icon'] as String?,
      sha: json['SHA'] as String?,
    );

Map<String, dynamic> _$ApplicationTemplateToJson(
        ApplicationTemplate instance) =>
    <String, dynamic>{
      'Appid': instance.appid,
      'Name': instance.name,
      'Website': instance.website,
      'License': instance.license,
      'Description': instance.description,
      'Enhanced': instance.enhanced,
      'tile_background': instance.tileBackground,
      'Icon': instance.icon,
      'SHA': instance.sha,
    };

AppList _$AppListFromJson(Map<String, dynamic> json) => AppList(
      appCount: json['AppCount'] as int?,
      apps: (json['Apps'] as List<dynamic>?)
              ?.map((e) =>
                  ApplicationTemplate.fromJson(e as Map<String, dynamic>))
              .toList() ??
          [],
    );

Map<String, dynamic> _$AppListToJson(AppList instance) => <String, dynamic>{
      'AppCount': instance.appCount,
      'Apps': instance.apps?.map((e) => e.toJson()).toList(),
    };
