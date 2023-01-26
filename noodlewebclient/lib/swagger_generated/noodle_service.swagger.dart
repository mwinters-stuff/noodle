// ignore_for_file: type=lint

import 'package:json_annotation/json_annotation.dart';
import 'package:collection/collection.dart';
import 'dart:convert';

import 'package:chopper/chopper.dart';

import 'client_mapping.dart';
import 'dart:async';
import 'package:chopper/chopper.dart' as chopper;

part 'noodle_service.swagger.chopper.dart';
part 'noodle_service.swagger.g.dart';

// **************************************************************************
// SwaggerChopperGenerator
// **************************************************************************

@ChopperApi()
abstract class NoodleService extends ChopperService {
  static NoodleService create({
    ChopperClient? client,
    Authenticator? authenticator,
    Uri? baseUrl,
    Iterable<dynamic>? interceptors,
  }) {
    if (client != null) {
      return _$NoodleService(client);
    }

    final newClient = ChopperClient(
        services: [_$NoodleService()],
        converter: $JsonSerializableConverter(),
        interceptors: interceptors ?? [],
        authenticator: authenticator,
        baseUrl: baseUrl ?? Uri.parse('http://localhost:8081/api'));
    return _$NoodleService(newClient);
  }

  ///Liveness check
  Future<chopper.Response<Object>> healthzGet() {
    return _healthzGet();
  }

  ///Liveness check
  @Get(path: '/healthz')
  Future<chopper.Response<Object>> _healthzGet();

  ///Readiness check
  Future<chopper.Response<Object>> readyzGet() {
    return _readyzGet();
  }

  ///Readiness check
  @Get(path: '/readyz')
  Future<chopper.Response<Object>> _readyzGet();

  ///
  ///@param userid
  Future<chopper.Response<List<User>>> noodleUsersGet({
    int? userid,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(User, () => User.fromJsonFactory);

    return _noodleUsersGet(userid: userid, remoteUser: remoteUser);
  }

  ///
  ///@param userid
  @Get(path: '/noodle/users')
  Future<chopper.Response<List<User>>> _noodleUsersGet({
    @Query('userid') int? userid,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param groupid
  Future<chopper.Response<List<Group>>> noodleGroupsGet({
    int? groupid,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(Group, () => Group.fromJsonFactory);

    return _noodleGroupsGet(groupid: groupid, remoteUser: remoteUser);
  }

  ///
  ///@param groupid
  @Get(path: '/noodle/groups')
  Future<chopper.Response<List<Group>>> _noodleGroupsGet({
    @Query('groupid') int? groupid,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param userid
  ///@param groupid
  Future<chopper.Response<List<UserGroup>>> noodleUserGroupsGet({
    int? userid,
    int? groupid,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(UserGroup, () => UserGroup.fromJsonFactory);

    return _noodleUserGroupsGet(userid: userid, groupid: groupid, remoteUser: remoteUser);
  }

  ///
  ///@param userid
  ///@param groupid
  @Get(path: '/noodle/user-groups')
  Future<chopper.Response<List<UserGroup>>> _noodleUserGroupsGet({
    @Query('userid') int? userid,
    @Query('groupid') int? groupid,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  Future<chopper.Response> noodleLdapReloadGet({String? remoteUser}) {
    return _noodleLdapReloadGet(remoteUser: remoteUser);
  }

  ///
  @Get(path: '/noodle/ldap/reload')
  Future<chopper.Response> _noodleLdapReloadGet({@Header('Remote-User') String? remoteUser});

  ///
  Future<chopper.Response> noodleHeimdallReloadGet({String? remoteUser}) {
    return _noodleHeimdallReloadGet(remoteUser: remoteUser);
  }

  ///
  @Get(path: '/noodle/heimdall/reload')
  Future<chopper.Response> _noodleHeimdallReloadGet({@Header('Remote-User') String? remoteUser});

  ///
  Future<chopper.Response<List<Tab>>> noodleTabsGet({String? remoteUser}) {
    generatedMapping.putIfAbsent(Tab, () => Tab.fromJsonFactory);

    return _noodleTabsGet(remoteUser: remoteUser);
  }

  ///
  @Get(path: '/noodle/tabs')
  Future<chopper.Response<List<Tab>>> _noodleTabsGet({@Header('Remote-User') String? remoteUser});

  ///
  ///@param action
  ///@param tab
  Future<chopper.Response<Tab>> noodleTabsPost({
    required String? action,
    required Tab? tab,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(Tab, () => Tab.fromJsonFactory);

    return _noodleTabsPost(action: action, tab: tab, remoteUser: remoteUser);
  }

  ///
  ///@param action
  ///@param tab
  @Post(path: '/noodle/tabs')
  Future<chopper.Response<Tab>> _noodleTabsPost({
    @Query('action') required String? action,
    @Body() required Tab? tab,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param tabid
  Future<chopper.Response> noodleTabsDelete({
    required int? tabid,
    String? remoteUser,
  }) {
    return _noodleTabsDelete(tabid: tabid, remoteUser: remoteUser);
  }

  ///
  ///@param tabid
  @Delete(path: '/noodle/tabs')
  Future<chopper.Response> _noodleTabsDelete({
    @Query('tabid') required int? tabid,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param tab_id
  Future<chopper.Response<List<ApplicationTab>>> noodleApplicationTabsGet({
    required int? tabId,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(ApplicationTab, () => ApplicationTab.fromJsonFactory);

    return _noodleApplicationTabsGet(tabId: tabId, remoteUser: remoteUser);
  }

  ///
  ///@param tab_id
  @Get(path: '/noodle/application-tabs')
  Future<chopper.Response<List<ApplicationTab>>> _noodleApplicationTabsGet({
    @Query('tab_id') required int? tabId,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param action
  ///@param application_tab
  Future<chopper.Response<ApplicationTab>> noodleApplicationTabsPost({
    required String? action,
    required ApplicationTab? applicationTab,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(ApplicationTab, () => ApplicationTab.fromJsonFactory);

    return _noodleApplicationTabsPost(action: action, applicationTab: applicationTab, remoteUser: remoteUser);
  }

  ///
  ///@param action
  ///@param application_tab
  @Post(path: '/noodle/application-tabs')
  Future<chopper.Response<ApplicationTab>> _noodleApplicationTabsPost({
    @Query('action') required String? action,
    @Body() required ApplicationTab? applicationTab,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param application_tab_id
  Future<chopper.Response> noodleApplicationTabsDelete({
    required int? applicationTabId,
    String? remoteUser,
  }) {
    return _noodleApplicationTabsDelete(applicationTabId: applicationTabId, remoteUser: remoteUser);
  }

  ///
  ///@param application_tab_id
  @Delete(path: '/noodle/application-tabs')
  Future<chopper.Response> _noodleApplicationTabsDelete({
    @Query('application_tab_id') required int? applicationTabId,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param user_id
  Future<chopper.Response<List<UserApplications>>> noodleUserApplicationsGet({
    required int? userId,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(UserApplications, () => UserApplications.fromJsonFactory);

    return _noodleUserApplicationsGet(userId: userId, remoteUser: remoteUser);
  }

  ///
  ///@param user_id
  @Get(path: '/noodle/user-applications')
  Future<chopper.Response<List<UserApplications>>> _noodleUserApplicationsGet({
    @Query('user_id') required int? userId,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param user_application
  Future<chopper.Response<UserApplications>> noodleUserApplicationsPost({
    required UserApplications? userApplication,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(UserApplications, () => UserApplications.fromJsonFactory);

    return _noodleUserApplicationsPost(userApplication: userApplication, remoteUser: remoteUser);
  }

  ///
  ///@param user_application
  @Post(path: '/noodle/user-applications')
  Future<chopper.Response<UserApplications>> _noodleUserApplicationsPost({
    @Body() required UserApplications? userApplication,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param user_application_id
  Future<chopper.Response> noodleUserApplicationsDelete({
    required int? userApplicationId,
    String? remoteUser,
  }) {
    return _noodleUserApplicationsDelete(userApplicationId: userApplicationId, remoteUser: remoteUser);
  }

  ///
  ///@param user_application_id
  @Delete(path: '/noodle/user-applications')
  Future<chopper.Response> _noodleUserApplicationsDelete({
    @Query('user_application_id') required int? userApplicationId,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param group_id
  Future<chopper.Response<List<GroupApplications>>> noodleGroupApplicationsGet({
    required int? groupId,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(GroupApplications, () => GroupApplications.fromJsonFactory);

    return _noodleGroupApplicationsGet(groupId: groupId, remoteUser: remoteUser);
  }

  ///
  ///@param group_id
  @Get(path: '/noodle/group-applications')
  Future<chopper.Response<List<GroupApplications>>> _noodleGroupApplicationsGet({
    @Query('group_id') required int? groupId,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param group_application
  Future<chopper.Response<GroupApplications>> noodleGroupApplicationsPost({
    required GroupApplications? groupApplication,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(GroupApplications, () => GroupApplications.fromJsonFactory);

    return _noodleGroupApplicationsPost(groupApplication: groupApplication, remoteUser: remoteUser);
  }

  ///
  ///@param group_application
  @Post(path: '/noodle/group-applications')
  Future<chopper.Response<GroupApplications>> _noodleGroupApplicationsPost({
    @Body() required GroupApplications? groupApplication,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param group_application_id
  Future<chopper.Response> noodleGroupApplicationsDelete({
    required int? groupApplicationId,
    String? remoteUser,
  }) {
    return _noodleGroupApplicationsDelete(groupApplicationId: groupApplicationId, remoteUser: remoteUser);
  }

  ///
  ///@param group_application_id
  @Delete(path: '/noodle/group-applications')
  Future<chopper.Response> _noodleGroupApplicationsDelete({
    @Query('group_application_id') required int? groupApplicationId,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param search
  Future<chopper.Response<List<ApplicationTemplate>>> noodleAppTemplatesGet({
    required String? search,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(ApplicationTemplate, () => ApplicationTemplate.fromJsonFactory);

    return _noodleAppTemplatesGet(search: search, remoteUser: remoteUser);
  }

  ///
  ///@param search
  @Get(path: '/noodle/app-templates')
  Future<chopper.Response<List<ApplicationTemplate>>> _noodleAppTemplatesGet({
    @Query('search') required String? search,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param application_id
  ///@param application_template
  Future<chopper.Response<List<Application>>> noodleApplicationsGet({
    int? applicationId,
    String? applicationTemplate,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(Application, () => Application.fromJsonFactory);

    return _noodleApplicationsGet(applicationId: applicationId, applicationTemplate: applicationTemplate, remoteUser: remoteUser);
  }

  ///
  ///@param application_id
  ///@param application_template
  @Get(path: '/noodle/applications')
  Future<chopper.Response<List<Application>>> _noodleApplicationsGet({
    @Query('application_id') int? applicationId,
    @Query('application_template') String? applicationTemplate,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param application
  Future<chopper.Response<Application>> noodleApplicationsPost({
    required Application? application,
    String? remoteUser,
  }) {
    generatedMapping.putIfAbsent(Application, () => Application.fromJsonFactory);

    return _noodleApplicationsPost(application: application, remoteUser: remoteUser);
  }

  ///
  ///@param application
  @Post(path: '/noodle/applications')
  Future<chopper.Response<Application>> _noodleApplicationsPost({
    @Body() required Application? application,
    @Header('Remote-User') String? remoteUser,
  });

  ///
  ///@param application_id
  Future<chopper.Response> noodleApplicationsDelete({
    required int? applicationId,
    String? remoteUser,
  }) {
    return _noodleApplicationsDelete(applicationId: applicationId, remoteUser: remoteUser);
  }

  ///
  ///@param application_id
  @Delete(path: '/noodle/applications')
  Future<chopper.Response> _noodleApplicationsDelete({
    @Query('application_id') required int? applicationId,
    @Header('Remote-User') String? remoteUser,
  });
}

@JsonSerializable(explicitToJson: true)
class Error {
  Error({
    this.code,
    this.message,
    this.fields,
  });

  factory Error.fromJson(Map<String, dynamic> json) => _$ErrorFromJson(json);

  @JsonKey(name: 'code')
  final num? code;
  @JsonKey(name: 'message')
  final String? message;
  @JsonKey(name: 'fields')
  final String? fields;
  static const fromJsonFactory = _$ErrorFromJson;
  static const toJsonFactory = _$ErrorToJson;
  Map<String, dynamic> toJson() => _$ErrorToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is Error &&
            (identical(other.code, code) || const DeepCollectionEquality().equals(other.code, code)) &&
            (identical(other.message, message) || const DeepCollectionEquality().equals(other.message, message)) &&
            (identical(other.fields, fields) || const DeepCollectionEquality().equals(other.fields, fields)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode => const DeepCollectionEquality().hash(code) ^ const DeepCollectionEquality().hash(message) ^ const DeepCollectionEquality().hash(fields) ^ runtimeType.hashCode;
}

extension $ErrorExtension on Error {
  Error copyWith({num? code, String? message, String? fields}) {
    return Error(code: code ?? this.code, message: message ?? this.message, fields: fields ?? this.fields);
  }

  Error copyWithWrapped({Wrapped<num?>? code, Wrapped<String?>? message, Wrapped<String?>? fields}) {
    return Error(code: (code != null ? code.value : this.code), message: (message != null ? message.value : this.message), fields: (fields != null ? fields.value : this.fields));
  }
}

@JsonSerializable(explicitToJson: true)
class User {
  User({
    this.id,
    this.username,
    this.dn,
    this.displayName,
    this.givenName,
    this.surname,
    this.uidNumber,
  });

  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);

  @JsonKey(name: 'Id')
  final num? id;
  @JsonKey(name: 'Username')
  final String? username;
  @JsonKey(name: 'DN')
  final String? dn;
  @JsonKey(name: 'DisplayName')
  final String? displayName;
  @JsonKey(name: 'GivenName')
  final String? givenName;
  @JsonKey(name: 'Surname')
  final String? surname;
  @JsonKey(name: 'UidNumber')
  final int? uidNumber;
  static const fromJsonFactory = _$UserFromJson;
  static const toJsonFactory = _$UserToJson;
  Map<String, dynamic> toJson() => _$UserToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is User &&
            (identical(other.id, id) || const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.username, username) || const DeepCollectionEquality().equals(other.username, username)) &&
            (identical(other.dn, dn) || const DeepCollectionEquality().equals(other.dn, dn)) &&
            (identical(other.displayName, displayName) || const DeepCollectionEquality().equals(other.displayName, displayName)) &&
            (identical(other.givenName, givenName) || const DeepCollectionEquality().equals(other.givenName, givenName)) &&
            (identical(other.surname, surname) || const DeepCollectionEquality().equals(other.surname, surname)) &&
            (identical(other.uidNumber, uidNumber) || const DeepCollectionEquality().equals(other.uidNumber, uidNumber)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode =>
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(username) ^
      const DeepCollectionEquality().hash(dn) ^
      const DeepCollectionEquality().hash(displayName) ^
      const DeepCollectionEquality().hash(givenName) ^
      const DeepCollectionEquality().hash(surname) ^
      const DeepCollectionEquality().hash(uidNumber) ^
      runtimeType.hashCode;
}

extension $UserExtension on User {
  User copyWith({num? id, String? username, String? dn, String? displayName, String? givenName, String? surname, int? uidNumber}) {
    return User(
        id: id ?? this.id,
        username: username ?? this.username,
        dn: dn ?? this.dn,
        displayName: displayName ?? this.displayName,
        givenName: givenName ?? this.givenName,
        surname: surname ?? this.surname,
        uidNumber: uidNumber ?? this.uidNumber);
  }

  User copyWithWrapped(
      {Wrapped<num?>? id, Wrapped<String?>? username, Wrapped<String?>? dn, Wrapped<String?>? displayName, Wrapped<String?>? givenName, Wrapped<String?>? surname, Wrapped<int?>? uidNumber}) {
    return User(
        id: (id != null ? id.value : this.id),
        username: (username != null ? username.value : this.username),
        dn: (dn != null ? dn.value : this.dn),
        displayName: (displayName != null ? displayName.value : this.displayName),
        givenName: (givenName != null ? givenName.value : this.givenName),
        surname: (surname != null ? surname.value : this.surname),
        uidNumber: (uidNumber != null ? uidNumber.value : this.uidNumber));
  }
}

@JsonSerializable(explicitToJson: true)
class Group {
  Group({
    this.id,
    this.dn,
    this.name,
  });

  factory Group.fromJson(Map<String, dynamic> json) => _$GroupFromJson(json);

  @JsonKey(name: 'Id')
  final num? id;
  @JsonKey(name: 'DN')
  final String? dn;
  @JsonKey(name: 'Name')
  final String? name;
  static const fromJsonFactory = _$GroupFromJson;
  static const toJsonFactory = _$GroupToJson;
  Map<String, dynamic> toJson() => _$GroupToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is Group &&
            (identical(other.id, id) || const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.dn, dn) || const DeepCollectionEquality().equals(other.dn, dn)) &&
            (identical(other.name, name) || const DeepCollectionEquality().equals(other.name, name)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode => const DeepCollectionEquality().hash(id) ^ const DeepCollectionEquality().hash(dn) ^ const DeepCollectionEquality().hash(name) ^ runtimeType.hashCode;
}

extension $GroupExtension on Group {
  Group copyWith({num? id, String? dn, String? name}) {
    return Group(id: id ?? this.id, dn: dn ?? this.dn, name: name ?? this.name);
  }

  Group copyWithWrapped({Wrapped<num?>? id, Wrapped<String?>? dn, Wrapped<String?>? name}) {
    return Group(id: (id != null ? id.value : this.id), dn: (dn != null ? dn.value : this.dn), name: (name != null ? name.value : this.name));
  }
}

@JsonSerializable(explicitToJson: true)
class UserGroup {
  UserGroup({
    this.id,
    this.groupId,
    this.groupDN,
    this.groupName,
    this.userId,
    this.userDN,
    this.userName,
  });

  factory UserGroup.fromJson(Map<String, dynamic> json) => _$UserGroupFromJson(json);

  @JsonKey(name: 'Id')
  final num? id;
  @JsonKey(name: 'GroupId')
  final num? groupId;
  @JsonKey(name: 'GroupDN')
  final String? groupDN;
  @JsonKey(name: 'GroupName')
  final String? groupName;
  @JsonKey(name: 'UserId')
  final num? userId;
  @JsonKey(name: 'UserDN')
  final String? userDN;
  @JsonKey(name: 'UserName')
  final String? userName;
  static const fromJsonFactory = _$UserGroupFromJson;
  static const toJsonFactory = _$UserGroupToJson;
  Map<String, dynamic> toJson() => _$UserGroupToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is UserGroup &&
            (identical(other.id, id) || const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.groupId, groupId) || const DeepCollectionEquality().equals(other.groupId, groupId)) &&
            (identical(other.groupDN, groupDN) || const DeepCollectionEquality().equals(other.groupDN, groupDN)) &&
            (identical(other.groupName, groupName) || const DeepCollectionEquality().equals(other.groupName, groupName)) &&
            (identical(other.userId, userId) || const DeepCollectionEquality().equals(other.userId, userId)) &&
            (identical(other.userDN, userDN) || const DeepCollectionEquality().equals(other.userDN, userDN)) &&
            (identical(other.userName, userName) || const DeepCollectionEquality().equals(other.userName, userName)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode =>
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(groupId) ^
      const DeepCollectionEquality().hash(groupDN) ^
      const DeepCollectionEquality().hash(groupName) ^
      const DeepCollectionEquality().hash(userId) ^
      const DeepCollectionEquality().hash(userDN) ^
      const DeepCollectionEquality().hash(userName) ^
      runtimeType.hashCode;
}

extension $UserGroupExtension on UserGroup {
  UserGroup copyWith({num? id, num? groupId, String? groupDN, String? groupName, num? userId, String? userDN, String? userName}) {
    return UserGroup(
        id: id ?? this.id,
        groupId: groupId ?? this.groupId,
        groupDN: groupDN ?? this.groupDN,
        groupName: groupName ?? this.groupName,
        userId: userId ?? this.userId,
        userDN: userDN ?? this.userDN,
        userName: userName ?? this.userName);
  }

  UserGroup copyWithWrapped(
      {Wrapped<num?>? id, Wrapped<num?>? groupId, Wrapped<String?>? groupDN, Wrapped<String?>? groupName, Wrapped<num?>? userId, Wrapped<String?>? userDN, Wrapped<String?>? userName}) {
    return UserGroup(
        id: (id != null ? id.value : this.id),
        groupId: (groupId != null ? groupId.value : this.groupId),
        groupDN: (groupDN != null ? groupDN.value : this.groupDN),
        groupName: (groupName != null ? groupName.value : this.groupName),
        userId: (userId != null ? userId.value : this.userId),
        userDN: (userDN != null ? userDN.value : this.userDN),
        userName: (userName != null ? userName.value : this.userName));
  }
}

@JsonSerializable(explicitToJson: true)
class UserApplications {
  UserApplications({
    this.id,
    this.applicationId,
    this.userId,
    this.application,
  });

  factory UserApplications.fromJson(Map<String, dynamic> json) => _$UserApplicationsFromJson(json);

  @JsonKey(name: 'Id')
  final num? id;
  @JsonKey(name: 'ApplicationId')
  final num? applicationId;
  @JsonKey(name: 'UserId')
  final num? userId;
  @JsonKey(name: 'Application')
  final Application? application;
  static const fromJsonFactory = _$UserApplicationsFromJson;
  static const toJsonFactory = _$UserApplicationsToJson;
  Map<String, dynamic> toJson() => _$UserApplicationsToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is UserApplications &&
            (identical(other.id, id) || const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.applicationId, applicationId) || const DeepCollectionEquality().equals(other.applicationId, applicationId)) &&
            (identical(other.userId, userId) || const DeepCollectionEquality().equals(other.userId, userId)) &&
            (identical(other.application, application) || const DeepCollectionEquality().equals(other.application, application)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode =>
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(applicationId) ^
      const DeepCollectionEquality().hash(userId) ^
      const DeepCollectionEquality().hash(application) ^
      runtimeType.hashCode;
}

extension $UserApplicationsExtension on UserApplications {
  UserApplications copyWith({num? id, num? applicationId, num? userId, Application? application}) {
    return UserApplications(id: id ?? this.id, applicationId: applicationId ?? this.applicationId, userId: userId ?? this.userId, application: application ?? this.application);
  }

  UserApplications copyWithWrapped({Wrapped<num?>? id, Wrapped<num?>? applicationId, Wrapped<num?>? userId, Wrapped<Application?>? application}) {
    return UserApplications(
        id: (id != null ? id.value : this.id),
        applicationId: (applicationId != null ? applicationId.value : this.applicationId),
        userId: (userId != null ? userId.value : this.userId),
        application: (application != null ? application.value : this.application));
  }
}

@JsonSerializable(explicitToJson: true)
class Tab {
  Tab({
    this.id,
    this.label,
    this.displayOrder,
  });

  factory Tab.fromJson(Map<String, dynamic> json) => _$TabFromJson(json);

  @JsonKey(name: 'Id')
  final num? id;
  @JsonKey(name: 'Label')
  final String? label;
  @JsonKey(name: 'DisplayOrder')
  final int? displayOrder;
  static const fromJsonFactory = _$TabFromJson;
  static const toJsonFactory = _$TabToJson;
  Map<String, dynamic> toJson() => _$TabToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is Tab &&
            (identical(other.id, id) || const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.label, label) || const DeepCollectionEquality().equals(other.label, label)) &&
            (identical(other.displayOrder, displayOrder) || const DeepCollectionEquality().equals(other.displayOrder, displayOrder)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode => const DeepCollectionEquality().hash(id) ^ const DeepCollectionEquality().hash(label) ^ const DeepCollectionEquality().hash(displayOrder) ^ runtimeType.hashCode;
}

extension $TabExtension on Tab {
  Tab copyWith({num? id, String? label, int? displayOrder}) {
    return Tab(id: id ?? this.id, label: label ?? this.label, displayOrder: displayOrder ?? this.displayOrder);
  }

  Tab copyWithWrapped({Wrapped<num?>? id, Wrapped<String?>? label, Wrapped<int?>? displayOrder}) {
    return Tab(id: (id != null ? id.value : this.id), label: (label != null ? label.value : this.label), displayOrder: (displayOrder != null ? displayOrder.value : this.displayOrder));
  }
}

@JsonSerializable(explicitToJson: true)
class GroupApplications {
  GroupApplications({
    this.id,
    this.applicationId,
    this.groupId,
    this.application,
  });

  factory GroupApplications.fromJson(Map<String, dynamic> json) => _$GroupApplicationsFromJson(json);

  @JsonKey(name: 'Id')
  final num? id;
  @JsonKey(name: 'ApplicationId')
  final num? applicationId;
  @JsonKey(name: 'GroupId')
  final num? groupId;
  @JsonKey(name: 'Application')
  final Application? application;
  static const fromJsonFactory = _$GroupApplicationsFromJson;
  static const toJsonFactory = _$GroupApplicationsToJson;
  Map<String, dynamic> toJson() => _$GroupApplicationsToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is GroupApplications &&
            (identical(other.id, id) || const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.applicationId, applicationId) || const DeepCollectionEquality().equals(other.applicationId, applicationId)) &&
            (identical(other.groupId, groupId) || const DeepCollectionEquality().equals(other.groupId, groupId)) &&
            (identical(other.application, application) || const DeepCollectionEquality().equals(other.application, application)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode =>
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(applicationId) ^
      const DeepCollectionEquality().hash(groupId) ^
      const DeepCollectionEquality().hash(application) ^
      runtimeType.hashCode;
}

extension $GroupApplicationsExtension on GroupApplications {
  GroupApplications copyWith({num? id, num? applicationId, num? groupId, Application? application}) {
    return GroupApplications(id: id ?? this.id, applicationId: applicationId ?? this.applicationId, groupId: groupId ?? this.groupId, application: application ?? this.application);
  }

  GroupApplications copyWithWrapped({Wrapped<num?>? id, Wrapped<num?>? applicationId, Wrapped<num?>? groupId, Wrapped<Application?>? application}) {
    return GroupApplications(
        id: (id != null ? id.value : this.id),
        applicationId: (applicationId != null ? applicationId.value : this.applicationId),
        groupId: (groupId != null ? groupId.value : this.groupId),
        application: (application != null ? application.value : this.application));
  }
}

@JsonSerializable(explicitToJson: true)
class Application {
  Application({
    this.id,
    this.templateAppid,
    this.name,
    this.website,
    this.license,
    this.description,
    this.enhanced,
    this.tileBackground,
    this.icon,
  });

  factory Application.fromJson(Map<String, dynamic> json) => _$ApplicationFromJson(json);

  @JsonKey(name: 'Id')
  final num? id;
  @JsonKey(name: 'TemplateAppid')
  final String? templateAppid;
  @JsonKey(name: 'Name')
  final String? name;
  @JsonKey(name: 'Website')
  final String? website;
  @JsonKey(name: 'License')
  final String? license;
  @JsonKey(name: 'Description')
  final String? description;
  @JsonKey(name: 'Enhanced')
  final bool? enhanced;
  @JsonKey(name: 'TileBackground')
  final String? tileBackground;
  @JsonKey(name: 'Icon')
  final String? icon;
  static const fromJsonFactory = _$ApplicationFromJson;
  static const toJsonFactory = _$ApplicationToJson;
  Map<String, dynamic> toJson() => _$ApplicationToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is Application &&
            (identical(other.id, id) || const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.templateAppid, templateAppid) || const DeepCollectionEquality().equals(other.templateAppid, templateAppid)) &&
            (identical(other.name, name) || const DeepCollectionEquality().equals(other.name, name)) &&
            (identical(other.website, website) || const DeepCollectionEquality().equals(other.website, website)) &&
            (identical(other.license, license) || const DeepCollectionEquality().equals(other.license, license)) &&
            (identical(other.description, description) || const DeepCollectionEquality().equals(other.description, description)) &&
            (identical(other.enhanced, enhanced) || const DeepCollectionEquality().equals(other.enhanced, enhanced)) &&
            (identical(other.tileBackground, tileBackground) || const DeepCollectionEquality().equals(other.tileBackground, tileBackground)) &&
            (identical(other.icon, icon) || const DeepCollectionEquality().equals(other.icon, icon)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode =>
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(templateAppid) ^
      const DeepCollectionEquality().hash(name) ^
      const DeepCollectionEquality().hash(website) ^
      const DeepCollectionEquality().hash(license) ^
      const DeepCollectionEquality().hash(description) ^
      const DeepCollectionEquality().hash(enhanced) ^
      const DeepCollectionEquality().hash(tileBackground) ^
      const DeepCollectionEquality().hash(icon) ^
      runtimeType.hashCode;
}

extension $ApplicationExtension on Application {
  Application copyWith({num? id, String? templateAppid, String? name, String? website, String? license, String? description, bool? enhanced, String? tileBackground, String? icon}) {
    return Application(
        id: id ?? this.id,
        templateAppid: templateAppid ?? this.templateAppid,
        name: name ?? this.name,
        website: website ?? this.website,
        license: license ?? this.license,
        description: description ?? this.description,
        enhanced: enhanced ?? this.enhanced,
        tileBackground: tileBackground ?? this.tileBackground,
        icon: icon ?? this.icon);
  }

  Application copyWithWrapped(
      {Wrapped<num?>? id,
      Wrapped<String?>? templateAppid,
      Wrapped<String?>? name,
      Wrapped<String?>? website,
      Wrapped<String?>? license,
      Wrapped<String?>? description,
      Wrapped<bool?>? enhanced,
      Wrapped<String?>? tileBackground,
      Wrapped<String?>? icon}) {
    return Application(
        id: (id != null ? id.value : this.id),
        templateAppid: (templateAppid != null ? templateAppid.value : this.templateAppid),
        name: (name != null ? name.value : this.name),
        website: (website != null ? website.value : this.website),
        license: (license != null ? license.value : this.license),
        description: (description != null ? description.value : this.description),
        enhanced: (enhanced != null ? enhanced.value : this.enhanced),
        tileBackground: (tileBackground != null ? tileBackground.value : this.tileBackground),
        icon: (icon != null ? icon.value : this.icon));
  }
}

@JsonSerializable(explicitToJson: true)
class ApplicationTab {
  ApplicationTab({
    this.id,
    this.applicationId,
    this.tabId,
    this.displayOrder,
    this.application,
  });

  factory ApplicationTab.fromJson(Map<String, dynamic> json) => _$ApplicationTabFromJson(json);

  @JsonKey(name: 'Id')
  final num? id;
  @JsonKey(name: 'ApplicationId')
  final num? applicationId;
  @JsonKey(name: 'TabId')
  final num? tabId;
  @JsonKey(name: 'DisplayOrder')
  final int? displayOrder;
  @JsonKey(name: 'Application')
  final Application? application;
  static const fromJsonFactory = _$ApplicationTabFromJson;
  static const toJsonFactory = _$ApplicationTabToJson;
  Map<String, dynamic> toJson() => _$ApplicationTabToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is ApplicationTab &&
            (identical(other.id, id) || const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.applicationId, applicationId) || const DeepCollectionEquality().equals(other.applicationId, applicationId)) &&
            (identical(other.tabId, tabId) || const DeepCollectionEquality().equals(other.tabId, tabId)) &&
            (identical(other.displayOrder, displayOrder) || const DeepCollectionEquality().equals(other.displayOrder, displayOrder)) &&
            (identical(other.application, application) || const DeepCollectionEquality().equals(other.application, application)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode =>
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(applicationId) ^
      const DeepCollectionEquality().hash(tabId) ^
      const DeepCollectionEquality().hash(displayOrder) ^
      const DeepCollectionEquality().hash(application) ^
      runtimeType.hashCode;
}

extension $ApplicationTabExtension on ApplicationTab {
  ApplicationTab copyWith({num? id, num? applicationId, num? tabId, int? displayOrder, Application? application}) {
    return ApplicationTab(
        id: id ?? this.id,
        applicationId: applicationId ?? this.applicationId,
        tabId: tabId ?? this.tabId,
        displayOrder: displayOrder ?? this.displayOrder,
        application: application ?? this.application);
  }

  ApplicationTab copyWithWrapped({Wrapped<num?>? id, Wrapped<num?>? applicationId, Wrapped<num?>? tabId, Wrapped<int?>? displayOrder, Wrapped<Application?>? application}) {
    return ApplicationTab(
        id: (id != null ? id.value : this.id),
        applicationId: (applicationId != null ? applicationId.value : this.applicationId),
        tabId: (tabId != null ? tabId.value : this.tabId),
        displayOrder: (displayOrder != null ? displayOrder.value : this.displayOrder),
        application: (application != null ? application.value : this.application));
  }
}

@JsonSerializable(explicitToJson: true)
class ApplicationTemplate {
  ApplicationTemplate({
    this.appid,
    this.name,
    this.website,
    this.license,
    this.description,
    this.enhanced,
    this.tileBackground,
    this.icon,
    this.sha,
  });

  factory ApplicationTemplate.fromJson(Map<String, dynamic> json) => _$ApplicationTemplateFromJson(json);

  @JsonKey(name: 'Appid')
  final String? appid;
  @JsonKey(name: 'Name')
  final String? name;
  @JsonKey(name: 'Website')
  final String? website;
  @JsonKey(name: 'License')
  final String? license;
  @JsonKey(name: 'Description')
  final String? description;
  @JsonKey(name: 'Enhanced')
  final bool? enhanced;
  @JsonKey(name: 'tile_background')
  final String? tileBackground;
  @JsonKey(name: 'Icon')
  final String? icon;
  @JsonKey(name: 'SHA')
  final String? sha;
  static const fromJsonFactory = _$ApplicationTemplateFromJson;
  static const toJsonFactory = _$ApplicationTemplateToJson;
  Map<String, dynamic> toJson() => _$ApplicationTemplateToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is ApplicationTemplate &&
            (identical(other.appid, appid) || const DeepCollectionEquality().equals(other.appid, appid)) &&
            (identical(other.name, name) || const DeepCollectionEquality().equals(other.name, name)) &&
            (identical(other.website, website) || const DeepCollectionEquality().equals(other.website, website)) &&
            (identical(other.license, license) || const DeepCollectionEquality().equals(other.license, license)) &&
            (identical(other.description, description) || const DeepCollectionEquality().equals(other.description, description)) &&
            (identical(other.enhanced, enhanced) || const DeepCollectionEquality().equals(other.enhanced, enhanced)) &&
            (identical(other.tileBackground, tileBackground) || const DeepCollectionEquality().equals(other.tileBackground, tileBackground)) &&
            (identical(other.icon, icon) || const DeepCollectionEquality().equals(other.icon, icon)) &&
            (identical(other.sha, sha) || const DeepCollectionEquality().equals(other.sha, sha)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode =>
      const DeepCollectionEquality().hash(appid) ^
      const DeepCollectionEquality().hash(name) ^
      const DeepCollectionEquality().hash(website) ^
      const DeepCollectionEquality().hash(license) ^
      const DeepCollectionEquality().hash(description) ^
      const DeepCollectionEquality().hash(enhanced) ^
      const DeepCollectionEquality().hash(tileBackground) ^
      const DeepCollectionEquality().hash(icon) ^
      const DeepCollectionEquality().hash(sha) ^
      runtimeType.hashCode;
}

extension $ApplicationTemplateExtension on ApplicationTemplate {
  ApplicationTemplate copyWith({String? appid, String? name, String? website, String? license, String? description, bool? enhanced, String? tileBackground, String? icon, String? sha}) {
    return ApplicationTemplate(
        appid: appid ?? this.appid,
        name: name ?? this.name,
        website: website ?? this.website,
        license: license ?? this.license,
        description: description ?? this.description,
        enhanced: enhanced ?? this.enhanced,
        tileBackground: tileBackground ?? this.tileBackground,
        icon: icon ?? this.icon,
        sha: sha ?? this.sha);
  }

  ApplicationTemplate copyWithWrapped(
      {Wrapped<String?>? appid,
      Wrapped<String?>? name,
      Wrapped<String?>? website,
      Wrapped<String?>? license,
      Wrapped<String?>? description,
      Wrapped<bool?>? enhanced,
      Wrapped<String?>? tileBackground,
      Wrapped<String?>? icon,
      Wrapped<String?>? sha}) {
    return ApplicationTemplate(
        appid: (appid != null ? appid.value : this.appid),
        name: (name != null ? name.value : this.name),
        website: (website != null ? website.value : this.website),
        license: (license != null ? license.value : this.license),
        description: (description != null ? description.value : this.description),
        enhanced: (enhanced != null ? enhanced.value : this.enhanced),
        tileBackground: (tileBackground != null ? tileBackground.value : this.tileBackground),
        icon: (icon != null ? icon.value : this.icon),
        sha: (sha != null ? sha.value : this.sha));
  }
}

@JsonSerializable(explicitToJson: true)
class AppList {
  AppList({
    this.appCount,
    this.apps,
  });

  factory AppList.fromJson(Map<String, dynamic> json) => _$AppListFromJson(json);

  @JsonKey(name: 'AppCount')
  final int? appCount;
  @JsonKey(name: 'Apps', defaultValue: <ApplicationTemplate>[])
  final List<ApplicationTemplate>? apps;
  static const fromJsonFactory = _$AppListFromJson;
  static const toJsonFactory = _$AppListToJson;
  Map<String, dynamic> toJson() => _$AppListToJson(this);

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is AppList &&
            (identical(other.appCount, appCount) || const DeepCollectionEquality().equals(other.appCount, appCount)) &&
            (identical(other.apps, apps) || const DeepCollectionEquality().equals(other.apps, apps)));
  }

  @override
  String toString() => jsonEncode(this);

  @override
  int get hashCode => const DeepCollectionEquality().hash(appCount) ^ const DeepCollectionEquality().hash(apps) ^ runtimeType.hashCode;
}

extension $AppListExtension on AppList {
  AppList copyWith({int? appCount, List<ApplicationTemplate>? apps}) {
    return AppList(appCount: appCount ?? this.appCount, apps: apps ?? this.apps);
  }

  AppList copyWithWrapped({Wrapped<int?>? appCount, Wrapped<List<ApplicationTemplate>?>? apps}) {
    return AppList(appCount: (appCount != null ? appCount.value : this.appCount), apps: (apps != null ? apps.value : this.apps));
  }
}

typedef $JsonFactory<T> = T Function(Map<String, dynamic> json);

class $CustomJsonDecoder {
  $CustomJsonDecoder(this.factories);

  final Map<Type, $JsonFactory> factories;

  dynamic decode<T>(dynamic entity) {
    if (entity is Iterable) {
      return _decodeList<T>(entity);
    }

    if (entity is T) {
      return entity;
    }

    if (isTypeOf<T, Map>()) {
      return entity;
    }

    if (isTypeOf<T, Iterable>()) {
      return entity;
    }

    if (entity is Map<String, dynamic>) {
      return _decodeMap<T>(entity);
    }

    return entity;
  }

  T _decodeMap<T>(Map<String, dynamic> values) {
    final jsonFactory = factories[T];
    if (jsonFactory == null || jsonFactory is! $JsonFactory<T>) {
      return throw "Could not find factory for type $T. Is '$T: $T.fromJsonFactory' included in the CustomJsonDecoder instance creation in bootstrapper.dart?";
    }

    return jsonFactory(values);
  }

  List<T> _decodeList<T>(Iterable values) => values.where((v) => v != null).map<T>((v) => decode<T>(v) as T).toList();
}

class $JsonSerializableConverter extends chopper.JsonConverter {
  @override
  FutureOr<chopper.Response<ResultType>> convertResponse<ResultType, Item>(chopper.Response response) async {
    if (response.bodyString.isEmpty) {
      // In rare cases, when let's say 204 (no content) is returned -
      // we cannot decode the missing json with the result type specified
      return chopper.Response(response.base, null, error: response.error);
    }

    final jsonRes = await super.convertResponse(response);
    return jsonRes.copyWith<ResultType>(body: $jsonDecoder.decode<Item>(jsonRes.body) as ResultType);
  }
}

final $jsonDecoder = $CustomJsonDecoder(generatedMapping);

// ignore: unused_element
String? _dateToJson(DateTime? date) {
  if (date == null) {
    return null;
  }

  final year = date.year.toString();
  final month = date.month < 10 ? '0${date.month}' : date.month.toString();
  final day = date.day < 10 ? '0${date.day}' : date.day.toString();

  return '$year-$month-$day';
}

class Wrapped<T> {
  final T value;
  const Wrapped.value(this.value);
}
