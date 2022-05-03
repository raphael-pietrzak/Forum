# Forum

## 1) Introduction
Nous devions créer un forum dans lequel des utilisateurs peuvent s'inscrirent pour poster des publications comme des photos ou commenter les posts de d'autres utilisateurs, mais ce n'est pas tout.

Nous allons donc vous présenter toutes les fonctionnalités de ce forum dans ce document.
<br>
<br>
## 2) Lancement du forum
### a) Localhost:5050
Pour lancer le forum, il faut rentrer la commande go run main/main.go dans le terminal et ensuite ouvrir le navigateur web au localhost 5500 en tapant "http://localhost:5500" dans la barre de recherche.
<br>
<br>
### b) Serveur web
Notre forum est également en ligne sur le net et accessible à l'adresse suivante : ...
<br>
<br>
## 3) Les fonctionnalités
Quand on accède au site, l'utilisateur n'est pas connecté alors il est considéré comment un "guest", il ne peut ni poster ni commenter le post de qui que ce soit. Il peut juste voir les posts. Il peut également les filtrer en fonction de leur date de publication de leur(s) catégorie(s) et de leur nombre de likes.
<br>
<br>
### a) Les Users
Quand un guest s'inscrit, il doit se connecter. A savoir que notre forum laisse la possibilité de se connecter avec ses identifiants google et permet même l'oublie de son mot de passe. Une fois connecter, le guest devient donc un "user". En plus des droits d'un guest, Il peut poster, commenter et il a un accès direct sur son profil.
Dans la page de profil, le user a une photo de profil prédéfini mais peut la changer par un gif, une image JPEG ou PNG. Il peut voir ses posts et combien il a de likes.
<br>
<br>
### b) Les modérateurs
Les modérateurs ont un peu plus de droits. Il peuvent supprimer des posts et peuvent également démote des users en guest si ces derniers ne respectent pas le forum.
<br>
<br>
### c) Les admins
Enfin, les admins qui eux ont les pleins pouvoirs. Ils ont en plus des modérateurs, la possibilité de promote tous les utilisateurs. Ils ont tous les droits.