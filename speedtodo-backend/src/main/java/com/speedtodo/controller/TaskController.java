package com.speedtodo.controller;
import com.speedtodo.model.Task;
import com.speedtodo.service.TaskService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.net.URI;
import java.util.List;
@RestController
@RequestMapping("/tasks")
public class TaskController {
    private final TaskService service;
    public TaskController(TaskService service) { this.service = service; }
    @GetMapping public List<Task> list(@RequestParam(required = false) String view) { return service.listTasks(view); }
    @GetMapping("/{id}") public Task get(@PathVariable Long id) { return service.getById(id); }
    @PostMapping public ResponseEntity<Task> create(@RequestBody Task task) { Task created = service.create(task); return ResponseEntity.created(URI.create("/tasks/" + created.getId())).body(created); }
    @PatchMapping("/{id}") public Task update(@PathVariable Long id, @RequestBody Task task) { return service.update(id, task); }
    @PostMapping("/{id}/complete") public Task complete(@PathVariable Long id) { return service.complete(id); }
    @PostMapping("/{id}/reopen") public Task reopen(@PathVariable Long id) { return service.reopen(id); }
    @DeleteMapping("/{id}") public ResponseEntity<Void> delete(@PathVariable Long id) { service.delete(id); return ResponseEntity.noContent().build(); }
}