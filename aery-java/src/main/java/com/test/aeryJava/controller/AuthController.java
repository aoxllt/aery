package com.test.aeryJava.controller;

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
        Cookie[] cookies = request.getCookies();
        if (cookies != null) {
            Optional<Cookie> existingCookie = Arrays.stream(cookies)
                    .filter(cookie -> "uuid".equals(cookie.getName()))
                    .findFirst();

            if (existingCookie.isPresent()) {
                String cookieValue = existingCookie.get().getValue();
                log.info("客户端带有的 uuid Cookie: " + cookieValue);
                return new ResponseEntity<>("欢迎回来，您的 uuid Cookie 是: " + cookieValue, HttpStatus.OK);
            }
        }

        String uuid = UUID.randomUUID().toString();
        Cookie sessionCookie = new Cookie("uuid", uuid);
        sessionCookie.setMaxAge(30 * 60); // 设置有效期
        sessionCookie.setDomain("10.22.72.209"); // 设置为访问者的域
        sessionCookie.setPath("/"); // 设置路径
        sessionCookie.setSecure(false); // 如果使用 HTTPS，请设置为 true
        sessionCookie.setHttpOnly(true); // 防止 JavaScript 访问

        log.info("创建了新的 uuid Cookie: " + uuid);
        response.addCookie(sessionCookie);

        return new ResponseEntity<>("欢迎访问，我们为您创建了新的 uuid Cookie: " + uuid, HttpStatus.OK);
    }

    @GetMapping("/get")
    public ResponseEntity<String> getCookie(HttpServletRequest request) {
        Cookie[] cookies = request.getCookies();
        if (cookies != null) {
            for (Cookie cookie : cookies) {
                if ("uuid".equals(cookie.getName())) {
                    log.info("akfdnoai;jdf;oiAJsfd;oiAJsfpoijassopidfjAPosjda" +
                            "flkjna水立方就API技术的【AjsfpijaspofdijAPOIdf" +
                            "aoifjoAIjf[poioAJsd[pojAsdpojapdofjk" +
                            "oaifjpoiAjsdf[pojapsdfjpajdf[poAjdsf" +
                            "aOUIfoiajdfpoisjdfposjakpofjAOwerdfr"+cookie.getName()+cookie.getValue());
                    return new ResponseEntity<>("接收到的 Cookie 值: " + cookie.getValue(), HttpStatus.OK);
                }
            }
        }
        log.info("没有找到");
        return new ResponseEntity<>("未找到 Cookie", HttpStatus.NOT_FOUND);
    }
}