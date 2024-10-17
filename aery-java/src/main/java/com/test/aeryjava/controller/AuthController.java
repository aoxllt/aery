package com.test.aeryjava.controller;

import jakarta.servlet.http.Cookie;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Arrays;
import java.util.Optional;
import java.util.UUID;

@RestController
@RequestMapping("/")
public class AuthController {
    private static final Logger log = LoggerFactory.getLogger(AuthController.class);

    @GetMapping("/")
    public ResponseEntity<String> visit(HttpServletRequest request, HttpServletResponse response) {
        // 检查 request.getCookies() 是否为 null
        Cookie[] cookies = request.getCookies();
        if (cookies != null) {
            // 先检查客户端请求是否携带了 session Cookie
            Optional<Cookie> existingCookie = Arrays.stream(cookies)
                    .filter(cookie -> "uuid".equals(cookie.getName()))
                    .findFirst();

            if (existingCookie.isPresent()) {
                // 如果客户端携带了 session Cookie，验证是否是合法的值
                String cookieValue = existingCookie.get().getValue();
                log.info("客户端带有的 uuid Cookie: " + cookieValue);
                return new ResponseEntity<>("欢迎回来，您的 uuid Cookie 是: " + cookieValue, HttpStatus.OK);
            }
        }

        // 如果客户端没有携带 session Cookie，则为它生成一个新的
        String uuid = UUID.randomUUID().toString();

        Cookie sessionCookie = new Cookie("uuid", uuid);
        sessionCookie.setHttpOnly(true); // 防止 JavaScript 访问 Cookie
        sessionCookie.setMaxAge(30 * 60); // 设置 Cookie 有效期为 30 分钟
        sessionCookie.setDomain("10.22.72.209"); // 设置为你的服务器IP
        sessionCookie.setPath("/"); // 设置路径为根路径

        log.info("创建了新的 uuid Cookie: " + uuid);
        response.addCookie(sessionCookie);

        return new ResponseEntity<>("欢迎访问，我们为您创建了新的 uuid Cookie: " + uuid, HttpStatus.OK);
    }
}