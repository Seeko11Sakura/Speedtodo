package com.speedtodo.controller;
import com.speedtodo.model.Tag;
import com.speedtodo.repository.TagRepository;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.net.URI;
import java.util.List;
@RestController
@RequestMapping("/tags")
public class TagController {
    private final TagRepository repository;
    public TagController(TagRepository repository) { this.repository = repository; }
    @GetMapping public List<Tag> list() { return repository.findAll(); }
    @PostMapping public ResponseEntity<Tag> create(@RequestBody Tag tag) { Tag created = repository.save(tag); return ResponseEntity.created(URI.create("/tags/" + created.getId())).body(created); }
}