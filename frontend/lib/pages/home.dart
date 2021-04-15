import 'package:flutter/material.dart';
import 'package:frontend/pages/components/sideBar.dart';

class HomeScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Home'),
      ),
      drawer: SideBar(),
      body: Center(
        child: ElevatedButton(
          child: Text('Launch screen'),
          onPressed: () {
            Navigator.pushNamed(context, '/library');
          },
        ),
      ),
    );
  }
}
