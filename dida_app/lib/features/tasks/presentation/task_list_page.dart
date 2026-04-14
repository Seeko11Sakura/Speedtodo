import 'package:flutter/material.dart';

class TaskListPage extends StatelessWidget {
  const TaskListPage({super.key, required this.view});

  final String view;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(view.toUpperCase())),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () {},
        label: const Text('Add Task'),
      ),
      body: const Center(child: Text('No tasks yet')),
    );
  }
}
