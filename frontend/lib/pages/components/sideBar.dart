import 'package:flutter/material.dart';

class SideBar extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: <Widget>[
          DrawerHeader(
            child: Text('Software Management'),
            decoration: BoxDecoration(
              color: Colors.blue,
            ),
          ),
          ListTile(
            title: Text('Home'),
            trailing: Icon(Icons.home),
            onTap: () {
              Navigator.pushNamed(context, '/');
            },
          ),
          ListTile(
            title: Text('Languages'),
            trailing: Icon(Icons.code),
            onTap: () {
              Navigator.pushNamed(context, '/library');
            },
          ),
          ListTile(
            title: Text('Person'),
            trailing: Icon(Icons.person),
            onTap: () {
              Navigator.pushNamed(context, '/person');
            },
          ),
          ListTile(
            title: Text('Settings'),
            trailing: Icon(Icons.settings),
            onTap: () {
              Navigator.pushNamed(context, '/');
            },
          ),
        ],
      ),
    );
  }
}
