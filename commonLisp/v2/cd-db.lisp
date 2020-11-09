(defvar *db* nil)

(defun make-cd (title artist rating ripped)
  (list :title title :artist artist :rating rating :ripped ripped))

(defun add-record (cd)
  (push cd *db*))

(add-record (make-cd "Welcome My Friend" "OKAMOTOS" 7 t))
(add-record (make-cd "Shekebon!" "bikke" 7 t))


(format t "------print db------")
*db*

