import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_form_builder/flutter_form_builder.dart';
import 'package:flutter_svg/svg.dart';
import 'package:form_builder_validators/form_builder_validators.dart';
import 'package:noodlewebclient/repositories/reositories.dart';

class MainScreen extends StatefulWidget {
  static const routeName = '/main';

  const MainScreen({Key? key}) : super(key: key);
  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  int _selectedIndex = 0;

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: <Widget>[
          NavigationRail(
            selectedIndex: _selectedIndex,
            groupAlignment: -1.0,
            labelType: NavigationRailLabelType.all,
            destinations: const <NavigationRailDestination>[
              NavigationRailDestination(
                icon: Icon(Icons.apps_outlined),
                selectedIcon: Icon(Icons.apps),
                label: Text('Apps'),
              ),
              NavigationRailDestination(
                icon: Icon(Icons.settings_outlined),
                selectedIcon: Icon(Icons.settings),
                label: Text('Settings'),
              ),
            ],
            onDestinationSelected: (value) {
              setState(() {
                _selectedIndex = value;
              });
            },
          ),
          const VerticalDivider(thickness: 1, width: 1),
          Expanded(
            child: Column(
              children: [
                SvgPicture.asset(
                  "images/noodle-icon.svg",
                  width: 256,
                  height: 256,
                ),
                Text(
                  "Noodle",
                  textAlign: TextAlign.center,
                  style: Theme.of(context).textTheme.displayLarge,
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
