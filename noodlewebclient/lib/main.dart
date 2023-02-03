import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:noodlewebclient/repositories/reositories.dart';
import 'package:noodlewebclient/widgets/widgets.dart';
import 'package:shared_preferences/shared_preferences.dart';

late SharedPreferences preferences;

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  SharedPreferences.getInstance().then((instance) {
    preferences = instance;
    runApp(const MyApp());
  });
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MultiRepositoryProvider(
      providers: [
        RepositoryProvider(create: (context) => SwaggerRepository("http://localhost:8081/api")),
      ],
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        title: 'Noodle',
        theme: ThemeData(
          useMaterial3: true,
          // This is the theme of your application.
          //
          // Try running your application with "flutter run". You'll see the
          // application has a blue toolbar. Then, without quitting the app, try
          // changing the primarySwatch below to Colors.green and then invoke
          // "hot reload" (press "r" in the console where you ran "flutter run",
          // or simply save your changes to "hot reload" in a Flutter IDE).
          // Notice that the counter didn't reset back to zero; the application
          // is not restarted.
          primarySwatch: Colors.grey,
        ),
        initialRoute: 'login',
        onGenerateRoute: AppRouter.generateRoute,
        // onGenerateRoute: (settings) {
        //   switch (settings.name) {
        //     case LoginScreen.routeName:
        //       return MaterialPageRoute(builder: (BuildContext context) => const LoginScreen());
        //     case MainScreen.routeName:
        //       return MaterialPageRoute(builder: (BuildContext context) => const MainScreen());
        //     default:
        //       return MaterialPageRoute(builder: (BuildContext context) => const LoginScreen());
        //   }
        // },
      ),
    );
  }
}
