import 'package:flutter/material.dart';

class TaskDetailPage extends StatelessWidget {
  const TaskDetailPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Task Detail')),
      body: const Padding(
        padding: EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            TextField(decoration: InputDecoration(labelText: 'Title')),
            SizedBox(height: 12),
            TextField(decoration: InputDecoration(labelText: 'Notes')),
          ],
        ),
      ),
    );
  }
}
