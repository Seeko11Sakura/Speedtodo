import '../domain/task.dart';

class TaskRepository {
  const TaskRepository();

  Future<List<Task>> fetchTasks(String view) async {
    return const <Task>[];
  }
}
